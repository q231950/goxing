/*
Package goxing allows using the XING platform from the command line.
It relies on xingapi for talking to the XING backend.
*/
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
			ShortName: "cs",
			Usage:     "Gets the contacts of the signed in user",
			Action:    app.LoadContactsAction,
		},
		{
			Name:      "contact",
			ShortName: "c",
			Usage:     "Gets the contact with the given name",
			Action:    app.LoadContactAction,
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
