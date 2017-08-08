package newrelic_api

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type deploymentRequest struct {
	Deployment DeploymentParams `json:"deployment"`
}

type DeploymentParams struct {
	Revision    string `json:"revision"`
	Changelog   string `json:"changelog"`
	Description string `json:"description"`
	User        string `json:"user"`
}

// Deployment sends a deployment tag to a New Relic application.
// https://docs.newrelic.com/docs/apm/new-relic-apm/maintenance/recording-deployments
func Deployment(n Client, id int64, dep DeploymentParams) error {
	d := deploymentRequest{dep}

	_, body, errs := gorequest.New().Post(fmt.Sprintf("https://api.newrelic.com/v2/applications/%d/deployments.json", id)).
		Set("X-Api-Key", n.key).
		Set("Content-Type", "application/json").
		Send(d).
		End()

	if len(errs) > 0 {
		return fmt.Errorf(body)
	}

	return nil
}
