# imaginary-exporter

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