// Package goxing allows using the XING platform from the command line
package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := NewApp(*cli.NewApp())
	app.Name = "XING cli"
	app.Usage = "xing"
	app.Author = "Martin Kim Dung-Pham"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:      "profile",
			ShortName: "p",
			Usage:     "Get my profile",
			Action:    app.LoadMeAction,
		},
		{
			Name:      "contacts",
			ShortName: "c",
			Usage:     "Get contacts for the given user id: c <userId>",
			Action:    app.LoadContactsAction,
		},
		{
			Name:      "messages",
			ShortName: "m",
			Usage:     "Get messages for the given user id: c <userId>",
			Action:    app.LoadMessagesAction,
		},
	}
	app.Run(os.Args)
}
