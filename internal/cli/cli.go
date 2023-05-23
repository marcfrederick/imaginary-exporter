package cli

import (
	"log"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/marcfrederick/imaginary-exporter/internal/cmd"
)

func Run(version string, args []string) {
	if err := NewApp(version).Run(args); err != nil {
		log.Fatalln(err)
	}
}

func NewApp(version string) *cli.App {
	return &cli.App{
		Name:     "imaginary-exporter",
		Version:  version,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{Name: "Marc Trölitzsch", Email: "Marc.Troelitzsch@gmail.com"},
		},
		Usage:                "Prometheus exporter for Imaginary metrics",
		EnableBashCompletion: true,

		Flags: cli.FlagsByName{
			&cli.StringFlag{
				Name:  "addr",
				Usage: "address to listen on",
				Value: ":8080",
			},
			&cli.StringFlag{
				Name:     "url",
				Usage:    "base url of the imaginary instance",
				Required: true,
			},
		},

		Action: func(cliCtx *cli.Context) error {
			exitCode, err := cmd.Run(cliCtx)
			if err != nil {
				return cli.Exit(err.Error(), exitCode)
			}
			return cli.Exit("", exitCode)
		},
	}
}
