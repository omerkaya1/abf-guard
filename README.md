# abf-guard

![Go](https://github.com/omerkaya1/abf-guard/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/omerkaya1/abf-guard)](https://goreportcard.com/report/github.com/omerkaya1/abf-guard)

ABF-Guard is a small service that prevents brute force login attacks.

## Usage
Full control of the project is available through the Makefile.

### Makefile targets and their description
- `setup`               - Install all the build and lint dependencies
- `mod`                 - Runs go mod on a project
- `fmt`                 - Runs goimports on all go files
- `test`                - Runs all unit tests
- `coverage`            - Runs all the tests and opens the coverage report
- `lint`                - Runs all the linters
- `vet`                 - Runs go vet
- `checks`              - Runs all checks for the project (go fmt, go lint, go vet)
- `build`               - Builds the project
- `run`                 - Runs the project in production mode
- `run-test`            - Runs the project for the local usage
- `gen`                 - Triggers code generation for the GRPC Server and Client API
- `gen-test`            - Triggers code generation for the GRPC Server and Client API for ITs
- `dockerbuild`         - Builds a docker image with the project
- `dockerpush`          - Publishes the docker image to the registry
- `docker-compose-up`   - Runs docker-compose command to kick-start the infrastructure
- `docker-compose-down` - Runs docker-compose command to turn down the infrastructure
- `integration`         - Runs the integration tests for the project
- `clean`               - Remove temporary files

## Client API
Service administration can be performed through the ABF-Guard CLI 
```
Run GRPC Web Service client for ABF-Guard

Usage:
  abf-guard grpc-client [command]

Examples:
  abf-guard grpc-client -h

Available Commands:
  add         Add an IP to a specified list
  auth        Authorisation request
  delete      Delete an IP from a specified list
  flush       Send a flush buckets request
  get         Get a list of IPs from a specified list
  purge       Purge single bucket

Flags:
  -b, --blacklist         blacklist or whitelist specification
  -e, --entity string     bucket name for removal
  -h, --help              help for grpc-client
  -s, --host string       host address (default "127.0.0.1")
  -i, --ip string         ip parameter
  -l, --login string      login parameter
  -w, --password string   password parameter
  -p, --port string       host port (default "6666")

Use "abf-guard grpc-client [command] --help" for more information about a command.
```

## ABF Guard server
```
Run GRPC Server for ABF-Guard

Usage:
  abf-guard grpc-server [flags]

Examples:
  abf-guard grpc-server -c /path/to/config.json

Flags:
  -c, --config string   -c, --config=/path/to/config.json
  -h, --help            help for grpc-server
```

## Settings
The app is primarily configured through a config.json file.
Settings for all the limitations for the bucket creation are also defined in this file.
Feel free to change them.

## Licence
This project is licenced under the GPL licence.
