package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := NewApp(*cli.NewApp())
	app.Name = "xing"
	app.Usage = "xing on the command line"
	app.Author = "Martin Kim Dung-Pham"
	app.Email = "kim.dung-pham@xing.com"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:      "me",
			ShortName: "p",
			Usage:     "Get my profile",
			Action: app.loadMeAction,
		},
		{
			Name:      "Contacts",
			ShortName: "c",
			Usage:     "Get contacts for the given user id: c <userId>",
			Action: app.LoadContactsAction,
		},
		{
			Name:      "Messages",
			ShortName: "m",
			Usage:     "Get messages for the given user id: c <userId>",
			Action: app.LoadMessagesAction,
		},
	}
	app.Run(os.Args)
}
