package main

import (
	"fmt"
	"os"

	"github.com/buckhx/denv/denvlib"
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
	app.Version = denvlib.Version
	app.Commands = []cli.Command{
		{
			Name:        "activate",
			Aliases:     []string{"a"},
			Usage:       "Activate an environment",
			Description: "Activate an environment by specifying it's name",
			Before:      argsRequired,
			Action: func(c *cli.Context) {
				env := c.Args().First()
				d, err := denvlib.Activate(env)
				if err != nil {
					fmt.Printf("%q does not exist", env)
				}
				fmt.Printf("Activated %s \n", d.Name())
			},
		},
		{
			Name:        "deactivate",
			Aliases:     []string{"d"},
			Usage:       "Deactivate the current environment",
			Description: "",
			Action: func(c *cli.Context) {
				undenv := denvlib.Deactivate()
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
				for d := range denvlib.List() {
					fmt.Println(d.Name())
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
				denvlib.Pull(remote, c.String("branch"))
				fmt.Printf("Pulled from %s successfully\n", remote)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "branch",
					Value: "master",
					Usage: "Branch to use from the remote",
				},
			},
		},
		{
			Name:        "push",
			Usage:       "Push up your current denvs to a remote server",
			Description: "Push your local devenv to the last server that was pulled",
			Before:      argsRequired,
			Action: func(c *cli.Context) {
				remote := c.Args().First()
				denvlib.Push(remote, c.String("branch"), c.String("message"))
				fmt.Printf("Pushed to %s successful\n", remote)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "branch",
					Value: "master",
					Usage: "Branch to use from the remote",
				},
				cli.StringFlag{
					Name:  "message",
					Value: "pushed",
					Usage: "Commit message for your push",
				},
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
				d, _ := denvlib.GetDenv(name)
				if d != nil && !c.Bool("force") {
					fmt.Printf("Denv %q already exists\nOverwrite with denv s -f %s\n", name, name)
				} else {
					fmt.Println("Snapshotting...")
					d := denvlib.Snapshot(name)
					included, _, _ := d.Files()
					for _, in := range included {
						fmt.Printf("\t%s\n", in)
					}
					fmt.Printf("Created a snapshot for %q\n", d.Name())
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
				d := denvlib.Which()
				// TODO maybe invert this logic?
				if d != nil {
					fmt.Println(d.Name())
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
	if !c.Args().Present() || c.Args().First() == "help" {
		cli.ShowSubcommandHelp(c)
		os.Exit(0)
	}
	return nil
}
