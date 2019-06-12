package commands

import (
	"github.com/fernandochristyanto/retsgo/app/core/commands/scaffold"
	"github.com/urfave/cli"
	"log"
	"os"
)

func RegisterCommands() {
	app := cli.NewApp()
	app.Name = "RETSgo"
	app.Description = "Scaffold and create a REST api with go in no time"
	app.Usage = "Scaffold and create a REST api with go in no time"
	app.Version = scaffold.GetVersion()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Fernando Christyanto",
			Email: "christyantofernando@gmail.com",
		},
	}
	app.UsageText = "retsgo [global options] command [command options] [args]"
	app.ArgsUsage = "[args]"

	app.Commands = []cli.Command{
		cli.Command{
			Name:        "dependencies",
			Usage:       "retsgo dependencies",
			Action:      PrintDependencies,
			Description: "List all dependencies needed to be installed",
		},
		cli.Command{
			Name:        "install",
			Usage:       "retsgo install",
			Action:      InstallDependencies,
			Description: "Installs external dependencies",
		},
		cli.Command{
			Name:        "newproject",
			Usage:       "retsgo newproject github.com/mygithubusername/mynewproject",
			ArgsUsage:   "repo",
			Action:      scaffold.BeginProjectScaffold,
			Description: "Generate a new project. retsgo newproject github.com/mygithubusername/mynewproject will create a new project under $GOPATH/src/github.com/mygithubusername/mynewproject",
		},
		cli.Command{
			Name:        "wiregen",
			Usage:       "retsgo wiregen",
			Action:      WireGen,
			Description: "Run this command in the root of the projectdir to generate wire-inject file in /projectdir/app/build/di/wire_gen.go. You should have a container.go file in /projectdir/app/startup/di/container.go",
		},
	}
	app.After = func(c *cli.Context) error {
		os.Exit(1)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
