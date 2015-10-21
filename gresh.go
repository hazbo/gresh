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
				leyra.get(c)
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
				fmt.Println("All configuration files were found!")
			},
		},
	}
	app.Run(os.Args)
}
