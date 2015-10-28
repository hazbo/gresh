package main

import (
	"fmt"
	"os"

	"gopkg.in/leyra/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "gresh"
	app.Usage = "a utility for creating a leyra app"
	app.Commands = []cli.Command{
		{
			Name:  "new",
			Usage: "create a new framework skeleton",
			Action: func(c *cli.Context) {
				if len(c.Args()) == 0 {
					fmt.Println("You must specify a name for your project")
					fmt.Println("e.g. gresh new my_app")
					return
				}
				leyra := new(leyra)
				leyra.get(c.Args())
				fmt.Println("")
				fmt.Println("Congrats! You now have a brand new copy of Leyra.")
				fmt.Println("The next thing you should do is to update your GOPATH to your newly created")
				fmt.Printf("project directory, %s.\n\n", c.Args()[0])
				fmt.Printf("This could look something like: export GOPATH=$(pwd)/%s\n\n", c.Args()[0])
				fmt.Printf("From now on, all gresh commands must be run from ./%s/src/leyra\n", c.Args()[0])
			},
		},
		{
			Name:  "configure",
			Usage: "configure your application",
			Action: func(c *cli.Context) {
				if findConfigFiles() != true {
					fmt.Println("No changes were made")
					return
				}
				// TODO: The following will output our new Makefile to stdout.
				// This will need to go somewhere around here, with all other
				// files generated from stubs.
				//
				// makefileFromStub()
				fmt.Println("All configuration files were found!")
			},
		},
	}
	app.Run(os.Args)
}
