package main

import (
	"fmt"
	"os"

	"github.com/buckhx/denv/api"
	"github.com/codegangsta/cli"
)

//go:generate go run scripts/include.go

func main() {
	client(os.Args)
}

func client(args []string) {
	app := cli.NewApp()
	app.Name = "denv"
	app.Usage = "Switch up your dev environments"
	app.Version = api.Version
	app.Commands = []cli.Command{
		{
			Name:        "activate",
			Aliases:     []string{"a"},
			Usage:       "Activate an environment",
			Description: "Activate an environment by specifying it's name",
			Before:      argsRequired,
			Action: func(c *cli.Context) {
				env := c.Args().First()
				denv, err := api.Activate(env)
				if err != nil {
					fmt.Printf("%q does not exist", env)
				}
				fmt.Printf("Activated %s \n", denv.Name())
			},
		},
		{
			Name:        "deactivate",
			Aliases:     []string{"d"},
			Usage:       "Deactivate the current environment",
			Description: "",
			Action: func(c *cli.Context) {
				undenv := api.Deactivate()
				if undenv != nil {
					fmt.Printf("Deactivated %s\n", undenv.Name())
				} else {
					fmt.Printf("No denv was active")
				}
			},
		},
		{
			Name:        "list",
			Aliases:     []string{"ls"},
			Usage:       "List the available environments",
			Description: "List the available environments",
			Action: func(c *cli.Context) {
				for denv := range api.List() {
					fmt.Println(denv.Name())
				}
			},
		},
		{
			Name:        "pull",
			Usage:       "Pull the denvs from a remote server",
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
			Usage:       "Push up your current denvs to a remote server",
			Description: "Push your local devenv to the last server that was pulled",
			Action: func(c *cli.Context) {
				out := api.Push()
				fmt.Println(out)
			},
		},
		{
			Name:        "snapshot",
			Aliases:     []string{"s", "save"},
			Usage:       "Snapshot home to new denv",
			Description: "Snapshot current home directory to Denv name",
			Before:      argsRequired,
			Action: func(c *cli.Context) {
				name := c.Args().First()
				denv, _ := api.GetDenv(name)
				if denv != nil && !c.Bool("force") {
					fmt.Printf("Denv %q already exists\nOverwrite with denv s -f %s\n", name, name)
				} else {
					fmt.Println("Snapshotting...")
					denv := api.Snapshot(name)
					included, _ := denv.Files()
					for _, in := range included {
						fmt.Printf("\t%s\n", in)
					}
					fmt.Printf("Created a snapshot for %q\n", denv.Name())
				}
			},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "force, f",
					Usage: "Force this command",
				},
			},
		},
		{
			Name:        "which",
			Aliases:     []string{"w"},
			Usage:       "Which denv is currently active",
			Description: "Which environemnt is currently activated",
			Action: func(c *cli.Context) {
				denv := api.Which()
				// TODO maybe invert this logic?
				if denv != nil {
					fmt.Println(denv.Name())
				} else {
					fmt.Println("No denv is active")
				}
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
