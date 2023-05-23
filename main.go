package main

import (
	"os"

	"github.com/marcfrederick/imaginary-exporter/internal/cli"
)

// Injected by GoReleaser during the release process
// https://goreleaser.com/cookbooks/using-main.version/
var version = "devel"

func main() {
	cli.Run(version, os.Args)
}
