package cmd

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"

	"github.com/marcfrederick/imaginary-exporter/pkg/collector"
	"github.com/marcfrederick/imaginary-exporter/pkg/imaginary"
)

func Run(cliCtx *cli.Context) (int, error) {
	address := cliCtx.String("addr")
	imaginaryURL := cliCtx.String("url")
	if !isURL(imaginaryURL) {
		return 1, fmt.Errorf("the given URL '%s' is invalid", imaginaryURL)
	}

	client := imaginary.NewClient(imaginaryURL)
	c := collector.NewImaginaryCollector(client)
	if err := prometheus.Register(c); err != nil {
		return 1, fmt.Errorf("error registering imaginary collector: %w", err)
	}

	http.Handle("/metrics", promhttp.Handler())
	log.Printf("listening on address %s", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		return 1, fmt.Errorf("error listening on address %s: %w", address, err)
	}

	return 0, nil
}

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
