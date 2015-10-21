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

				f := new(file)
				f.download("https://github.com/leyra/leyra/archive/master.zip")
				f.unzip("master.zip", c.Args()[0])
			},
		},
	}
	app.Run(os.Args)
}
