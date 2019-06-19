# MSV - Minimal HTTP Server in Go

[![Build Status](https://travis-ci.org/ribice/msv.svg?branch=master)](https://travis-ci.org/ribice/msv)
[![codecov](https://codecov.io/gh/ribice/msv/branch/master/graph/badge.svg)](https://codecov.io/gh/ribice/msv)
[![Go Report Card](https://goreportcard.com/badge/github.com/ribice/msv)](https://goreportcard.com/report/github.com/ribice/msv)
[![Maintainability](https://api.codeclimate.com/v1/badges/c3cb09dbc0bc43186464/maintainability)](https://codeclimate.com/github/ribice/msv/maintainability)

MSV is a minimal server implementation in Go. It wraps Gorilla multiplexer, lifecycle control, graceful shutdown and common middleware into a single reusable package.

## Introduction

Whenever I start working on a new project, the packages implemented in msv are the ones I always tend to write from scratch. The idea for msv is to prevent that, and keep the common server implementation reusable for multiple projects.

## What's included

MSV includes the following packages:

- MSV (Server): Basic server implementation that wraps `gorilla/mux` with net/http server. Includes graceful shutdown and Start methods on server.
- Bind: Marshalls and validates HTTP requests. `Binder` interface needs to be implemented in order for a struct to be passed to Bind.
- Render: Writes HTTP responses and status codes to clients.
- Middleware: Includes various http middlewares:
    - bauth: HTTP Basic Authentication
    - httplog: HTTP Request/Response logging
    - recocery: HTTP panics recoverer

## License

MSV is licensed under the MIT license. Check the [LICENSE](LICENSE.md) file for details.

## Author

[Emir Ribic](https://ribice.ba)