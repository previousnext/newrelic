package newrelic_api

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

// Applications is a collection of New Relic applications.
type Applications struct {
	Applications []Application `json:"applications"`
}

// Application is a New Relic application.
type Application struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Language     string `json:"language"`
	HealthStatus string `json:"health_status"`
}

// NameToApplicationID returns an App ID based on App Name.
func NameToApplicationID(n Client, name string) (int64, error) {
	resp, err := ListApplications(n)
	if err != nil {
		return 0, err
	}

	for _, app := range resp.Applications {
		if app.Name == name {
			return app.ID, nil
		}
	}

	return 0, fmt.Errorf("Cannot find application with name: %s", name)
}

// ListApplications returns a list of applications.
func ListApplications(n Client) (Applications, error) {
	var apps Applications

	_, _, errs := gorequest.New().Get("https://api.newrelic.com/v2/applications.json").
		Set("X-Api-Key", n.key).
		Set("Content-Type", "application/json").
		EndStruct(&apps)

	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Println(e)
		}

		return apps, errs[0]
	}

	return apps, nil
}
