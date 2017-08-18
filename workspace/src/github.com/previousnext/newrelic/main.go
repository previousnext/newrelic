package main

import (
	"fmt"
	"github.com/previousnext/go-newrelic"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app    = kingpin.New("newrelic", "A command-line interface for New Relic API.")
	apiKey = app.Flag("api-key", "New Relic API key. Use $NR_API_KEY environment var.").Envar("NR_API_KEY").Required().String()

	deploy            = app.Command("deployment", "Record a deployment.")
	deployAppName     = deploy.Flag("app", "Name of application in New Relic.").Required().String()
	deployRevision    = deploy.Flag("revision", "A unique ID for this deployment. Can be any string, but is usually a version number or a Git checksum.").Required().String()
	deployChangelog   = deploy.Flag("changelog", "A summary of what changed in this deployment, visible in the Deployments page when you select (selected deployment) > Change log.").String()
	deployDescription = deploy.Flag("description", "A high-level description of this deployment, visible in the Overview page and on the Deployments page when you select an individual deployment.").String()
	deployUser        = deploy.Flag("user", "A username to associate with the deployment, visible in the Overview page and on the Deployments page.").String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	// Record deployment command.
	case deploy.FullCommand():
		n := newrelic.New(*apiKey)

		deployCommand(n)
	}
}

func deployCommand(n newrelic.Client) {
	// Figure out the appID from the appName.
	appID, err := n.NameToApplicationID(*deployAppName)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error retrieving appID for '%s'", *deployAppName))
		panic(err)
	}

	dply := newrelic.DeploymentInput{
		Deployment: newrelic.Deployment{
			Revision:    *deployRevision,
			Changelog:   *deployChangelog,
			Description: *deployDescription,
			User:        *deployUser,
		},
	}

	e := n.Deployment(appID, dply)
	if e != nil {
		panic(e)
	}

	fmt.Println(fmt.Sprintf("Successfully tagged %s with version '%s'", *deployAppName, *deployRevision))
}
