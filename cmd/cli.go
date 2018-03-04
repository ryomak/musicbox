package cmd

import (
	"github.com/ryomak/musicbox/cmd/action"
	"github.com/urfave/cli"
)

func New(name, usage, version string) *cli.App {
	app := cli.NewApp()
	app.Name = name
	app.Usage = usage
	app.Version = version
	app.Commands = getCommands()
	return app
}

func getCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "show music list",
			Action: action.List,
		},
		{
			Name:   "play",
			Usage:  "play music",
			Action: action.Play,
		},
	}
}
