New Relic CLI
=======================

[![CircleCI](https://circleci.com/gh/previousnext/newrelic-cli.svg?style=svg)](https://circleci.com/gh/previousnext/newrelic-cli)

**Maintainer**: Nick Santamaria <nick.santamaria@previousnext.com.au>

This is a basic CLI utility for interfacing with the New Relic API. Currently this supports:

- [Deployments](https://docs.newrelic.com/docs/apm/new-relic-apm/maintenance/recording-deployments)

## Usage

Record a deployment in New Relic using the `deployment` command. It is recommended to provide the API key via the `NR_API_KEY` environment variable.

```sh
newrelic deploy --app="My App - Dev" --revision="v1.0.4"
``` 

View the inbuilt help text using the `--help` flag.
```sh
newrelic deployment --app=APP --revision=REVISION [<flags>]

Record a deployment.

Flags:
  --help                     Show context-sensitive help (also try --help-long
                             and --help-man).
  --api-key=API-KEY          New Relic API key. Use $NR_API_KEY environment var.
  --app=APP                  Name of application in New Relic.
  --revision=REVISION        A unique ID for this deployment. Can be any string,
                             but is usually a version number or a Git checksum.
  --changelog=CHANGELOG      A summary of what changed in this deployment,
                             visible in the Deployments page when you select
                             (selected deployment) > Change log.
  --description=DESCRIPTION  A high-level description of this deployment,
                             visible in the Overview page and on the Deployments
                             page when you select an individual deployment.
  --user=USER                A username to associate with the deployment,
                             visible in the Overview page and on the Deployments
                             page.
```

## Resources

- [New Relic - Recording Deployments](https://docs.newrelic.com/docs/apm/new-relic-apm/maintenance/recording-deployments)
- [Dave Cheney - Reproducible Builds](https://www.youtube.com/watch?v=c3dW80eO88I)

## Development

### Principles

* Code lives in the `workspace` directory

### Tools

* **Dependency management** - https://getgb.io
* **Build** - https://github.com/mitchellh/gox
* **Linting** - https://github.com/golang/lint

### Workflow

(While in the `workspace` directory)

**Installing a new dependency**

```bash
gb vendor fetch github.com/foo/bar
```

**Running quality checks**

```bash
make lint test
```

**Building binaries**

```bash
make build
```
