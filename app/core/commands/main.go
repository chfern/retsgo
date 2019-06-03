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
			Name:   "dependencies",
			Usage:  "List all dependencies needed to be installed",
			Action: PrintDependencies,
		},
		cli.Command{
			Name:   "install",
			Usage:  "Installs external dependencies",
			Action: InstallDependencies,
		},
		cli.Command{
			Name:        "newproject",
			Usage:       "Generate a new project",
			ArgsUsage:   "repo",
			Action:      scaffold.BeginProjectScaffold,
			Description: "retsgo newproject github.com/mygithubusername/mynewproject will create a new project under $GOPATH/src/github.com/mygithubusername/mynewproject",
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
