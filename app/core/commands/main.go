package commands

import (
	"fmt"
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
	app.Version = "1.0.0"
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
		{
			Name:    "template",
			Aliases: []string{"t"},
			Usage:   "options for task templates",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
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
