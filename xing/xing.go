package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := XINGApp{*cli.NewApp()}
	app.Name = "xing"
	app.Usage = "xing on the command line"
	app.Author = "Martin Kim Dung-Pham"
	app.Email = "kim.dung-pham@xing.com"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:      "me",
			ShortName: "m",
			Usage:     "Get my profile",
			Action: app.loadMeAction,
		},
		{
			Name:      "Contacts",
			ShortName: "c",
			Usage:     "Get contacts for the given user id: c <userId>",
			Action: app.loadUserAction,
		},
	}
	app.Run(os.Args)
}
