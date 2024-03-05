FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .

FROM scratch

WORKDIR /app

COPY --from=builder /app/imaginary-exporter /usr/bin/

ENTRYPOINT ["/usr/bin/imaginary-exporter"]
