# amctl

## Description

`amctl` is a command line tool for interacting with Alertmanager API.

It has been built using `go-swagger generate cli` command based on v2 Alertmanager OpenAPI specification.

## Installation

```bash
go install github.com/geekxflood/amctl@latest
```

## Usage:

`amctl` `command` `flags`

## Available Commands:

- `alert`
- `alertgroup`
- `completion`  Generate completion script
- `general`
- `help`        Help about any command
- `receiver`
- `silence`

### Flags:

- `--base-path` `string`   For example: `/api/v2/` (default `/api/v2/`)
- `--config` `string`      config file path
- `--debug`              output debug logs
- `--dry-run`           do not send the request to server
- `-h`, `--help`          help for `amctl`
- `--hostname` `string`   hostname of the service (default "localhost")
- `--scheme` `string`     Choose from: `http` (default `http`)

Use "`amctl` `command` `--help`" for more information about a command.
