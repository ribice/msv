# MSV - Minimal HTTP Server in Go

MSV is a minimal server implementation in Go. It wraps Gorilla multiplexer, lifecycle control, graceful shutdown and common middleware into a single reusable package.

## Why MSV

Whenever I start working on a new project, the packages implemented in msv are the ones I always tend to write from scratch. The idea for msv is to prevent that, and keep the common server implementation reusable for multiple projects.

## What's included

MSV includes the following packages:

### Server

### Bind

### Render

### Middleware

- bauth: 