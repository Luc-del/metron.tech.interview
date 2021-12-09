package application

import (
	"context"

	"github.com/urfave/cli/v2"
)

var (
	port = "port"
)

func BuildCLIApp(appVersion string) *cli.App {
	app := &cli.App{
		Name:   "remedy",
		Usage:  "series operation orchestration service",
		Action: makeRunInCliContext(appVersion),
	}

	app.Flags = []cli.Flag{
		&cli.IntFlag{
			Name:    port,
			EnvVars: []string{"HTTP_PORT"},
			Usage:   "Http port",
			Value:   6666,
		},
	}

	return app
}

func makeRunInCliContext(version string) cli.ActionFunc {
	return func(c *cli.Context) error {
		app, err := New(context.Background(), version)
		if err != nil {
			return err
		}

		return app.Run(c.Int(port))
	}
}
