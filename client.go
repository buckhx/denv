package main

import (
	"fmt"
	"os"

	"github.com/buckhx/devenv/api"
	"github.com/codegangsta/cli"
)

func main() {
	client(os.Args)
}

func client(args []string) {
	app := cli.NewApp()
	app.Name = "devenv"
	app.Usage = "Switch up your dev environments"
	app.Action = func(c *cli.Context) {
		if !c.Args().Present() {
			cli.ShowAppHelp(c)
			os.Exit(0)
		}
		in := c.Args().First()
		out := api.Echo(in)
		fmt.Println(out)
	}
	app.Commands = []cli.Command{
		{
			Name:  "upper",
			Usage: "Raise the steaks",
			Action: func(c *cli.Context) {
				in := c.Args().First()
				out := api.Upper(in)
				fmt.Println(out)
			},
		},
		{
			Name:  "lower",
			Usage: "Lower the bar",
			Action: func(c *cli.Context) {
				in := c.Args().First()
				out := api.Lower(in)
				fmt.Println(out)
			},
		},
	}
	app.Run(args)
}
