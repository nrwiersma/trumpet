package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/hamba/cmd/v2"
	"github.com/hamba/cmd/v2/term"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli/v2"
)

const flagService = "services"

var version = "¯\\_(ツ)_/¯"

func commands(ui term.Term) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "server",
			Usage: "Run the ren HTTP server",
			Flags: cmd.Flags{
				&cli.StringFlag{
					Name:     flagService,
					Usage:    "The path to the services config file.",
					EnvVars:  []string{"SERVICES"},
					Required: true,
				},
			}.Merge(cmd.LogFlags),
			Action: runServer(ui),
		},
	}
}

func main() {
	ui := newTerm()

	app := &cli.App{
		Name:     "trumpet",
		Version:  version,
		Commands: commands(ui),
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := app.RunContext(ctx, os.Args); err != nil {
		ui.Error(err.Error())
	}
}

func newTerm() term.Term {
	return term.Prefixed{
		ErrorPrefix: "Error: ",
		Term: term.Colored{
			ErrorColor: term.Red,
			Term: term.Basic{
				Writer:      os.Stdout,
				ErrorWriter: os.Stderr,
				Verbose:     false,
			},
		},
	}
}
