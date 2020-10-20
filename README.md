# [dotenv](https://git.rootprojects.org/root/dotenv)

A cross-platform tool to run a command with environment variables loaded from a given .env file.

**Example**:

```bash
dotenv -f .env -- node server.js
```

## Install

**Mac**, **Linux**:

```bash
curl -sS https://webinstall.dev/dotenv | bash
```

**Windows 10**:

```pwsh
curl -A MS https://webinstall.dev/dotenv | powershell
```

## Usage

```bash
dotenv v1.0.0 (17c7677) 2020-10-19T23:43:57Z

Usage:
    dotenv [-f .env.alternate] -- <command> [arguments...]

  -f string
    	path to .env file (default ".env")
  --help
    	print usage and exit
  --version
    	print version and exit

Example:
    dotenv -f .env -- caddy run --config Caddyfile
```

## Build

`dotenv` is written in Go, using [github.com/joho/godotenv](https://github.com/joho/godotenv).

Install Go:

```bash
curl -sS https://webinstall.dev/golang | bash
```

(see https://webinstall.dev/golang for details)

### Manually

```bash
export GOFLAGS="-mod=vendor"
go generate -mod=vendor ./...
go run -mod=vendor git.rootprojects.org/root/go-gitver/v2
go build -mod=vendor .
```

```bash
./dotenv --version
```

### With GoReleaser

Install GoReleaser:

```bash
curl -sS https://webinstall.dev/goreleaser | bash
```

```bash
goreleaser --snapshot --skip-publish --rm-dist
```

```bash
./dist/dotenv_darwin_amd64/dotenv --version
```

## License

Copyright 2020 The dotenv Authors.

This Source Code Form is subject to the terms of the Mozilla Public \
License, v. 2.0. If a copy of the MPL was not distributed with this \
file, You can obtain one at http://mozilla.org/MPL/2.0/.
