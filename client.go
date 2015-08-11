package main

import (
	"fmt"
	"os"

	"github.com/buckhx/denv/api"
	"github.com/codegangsta/cli"
)

func main() {
	client(os.Args)
}

func client(args []string) {
	app := cli.NewApp()
	app.Name = "devenv"
	app.Usage = "Switch up your dev environments"
	app.Commands = []cli.Command{
		{
			Name:        "activate",
			Aliases:     []string{"a"},
			Usage:       "Activate an environment",
			Description: "Activate an environment by specifying it's name",
			Before:      argsRequired,
			Action: func(c *cli.Context) {
				env := c.Args().First()
				api.Activate(env)
				fmt.Printf("Activated %s \n", env)
			},
		},
		{
			Name:        "list",
			Aliases:     []string{"ls"},
			Usage:       "devenv ls",
			Description: "List the available environments",
			Action: func(c *cli.Context) {
				out := api.List()
				fmt.Println(out)
			},
		},
		{
			Name:        "pull",
			Usage:       "devenv pull http://github.com/buckhx/denv",
			Description: "Pull from the remote devenv",
			Before:      argsRequired,
			Action: func(c *cli.Context) {
				remote := c.Args().First()
				out := api.Pull(remote)
				fmt.Println(out)
			},
		},
		{
			Name:        "push",
			Usage:       "devenv push",
			Description: "Push your local devenv to the last server that was pulled",
			Action: func(c *cli.Context) {
				out := api.Push()
				fmt.Println(out)
			},
		},
		{
			Name:        "which",
			Aliases:     []string{"w"},
			Usage:       "devenv w",
			Description: "Which environemnt is currently activated",
			Action: func(c *cli.Context) {
				out := api.Which()
				fmt.Println(out)
			},
		},
	}
	app.Run(args)
}

// Check to see if the are enough args
func argsRequired(c *cli.Context) error {
	if !c.Args().Present() {
		cli.ShowSubcommandHelp(c)
		os.Exit(0)
	}
	return nil
}
