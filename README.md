# imaginary-exporter

[![License: GPL-2.0](https://img.shields.io/badge/License-GPL--2.0-blue.svg)](https://opensource.org/licenses/GPL-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/marcfrederick/imaginary-exporter)](https://goreportcard.com/report/github.com/marcfrederick/imaginary-exporter)
[![GoDoc](https://pkg.go.dev/badge/github.com/marcfrederick/imaginary-exporter.svg)](https://pkg.go.dev/github.com/marcfrederick/imaginary-exporter)
[![Go Version](https://img.shields.io/badge/go%20version-1.20+-blue.svg)](https://golang.org/)

<div align="center">
    <img src="assets/logo.png" alt="imaginary-exporter logo" width="256px"/>
</div>

Prometheus exporter for [Imaginary](https://github.com/h2non/imaginary) written in Go.

## Compatibility

The exporter in compatible with all versions exporting the `/health` endpoint (v0.1.18+)
However, some metrics are only on versions v1.0.17+.

## Usage

```
-fetch-rate duration
    interval in milliseconds in which to fetch metrics (default 15s)
-url string
    url of the imaginary instance
```
