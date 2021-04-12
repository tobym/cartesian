# Cartesian

[![Build](https://github.com/leozz37/cartesian/actions/workflows/build.yml/badge.svg)](https://github.com/leozz37/cartesian/actions/workflows/build.yml)
[![Unit Tests](https://github.com/leozz37/cartesian/actions/workflows/unit_tests.yml/badge.svg)](https://github.com/leozz37/cartesian/actions/workflows/unit_tests.yml)
[![Docker](https://github.com/leozz37/cartesian/actions/workflows/docker.yml/badge.svg)](https://github.com/leozz37/cartesian/actions/workflows/docker.yml)
[![Docker Compose](https://github.com/leozz37/cartesian/actions/workflows/docker_compose.yml/badge.svg)](https://github.com/leozz37/cartesian/actions/workflows/docker_compose.yml)
[![Terraform](https://github.com/leozz37/cartesian/actions/workflows/terraform.yml/badge.svg)](https://github.com/leozz37/cartesian/actions/workflows/terraform.yml)

This is an API for Civi test.

## Contents

- [Quick-start](#quick-start)
- [Building](#building)
  - [Binary](#binary)
  - [Makefile](#makefile)
  - [Docker](#docker)
  - [Docker-compose](#docker-compose)
- [Testing](#testing)
  - [Unit Tests](#unit-tests)

## Quick-start

This is an REST API, made with Golang and Gin. You can manually run it or use docker-compose (recommended).

To install the API dependencies run:

```shell
$ go mod download
```

## Building

By far, the easiest way to get everything running is with `docker-compose`. See the [docker-compose](#docker-compose) section.

### Binary

To build the binary, run the following:

```shell
$ go build -o cartesian
```

To run the binary, run the following:

```shell
$ ./cartesian
```

Or simply:

```shell
$ go run main.go
```

### Makefile

To run the through the Makefile, run the following:

```shell
$ make run
```

### Docker

Make sure you have [Docker](https://www.docker.com/get-started) installed on your machine.

To build the Docker image, run the following:

```shell
$ docker build . -t cartesian
```

To run the Docker image, set the `$PORT` variable and run:

```shell
$ export PORT=8080

$ docker run --env "PORT=$PORT" -p $PORT:$PORT cartesian
```

### Docker-compose

To run the docker-compose, set the `$PORT` variable and run:

```shell
$ export PORT=8080

$ docker-compose up
```

## Testing

The unit testes are written with the default testing tool of Golang.

### Unit Tests

To run the unit tests, do the following:

```shell
$ go test -v ./...
```

To run the tests with coverage, do the following:

```shell
$ go test -v -covermode=count ./...
````
