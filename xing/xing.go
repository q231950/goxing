package main

import (
	"github.com/codegangsta/cli"
	"github.com/str1ngs/ansi/color"
	"os"
	"xingapi"
	"fmt"
	// "bufio"
)

func main() {
	app := cli.NewApp()
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
			Action: func(c *cli.Context) {

				client := new(xingapi.Client)	
				client.Me(func(me xingapi.User) {
						xingapi.PrintUser(me)
				})
			},
		},
		{
			Name:      "Contacts",
			ShortName: "c",
			Usage:     "Get contacts for the given user id: c <userid>",
			Action: func(c *cli.Context) {
				userid := c.Args().First()
				client := new(xingapi.Client)
				client.ContactsList(userid, func(list xingapi.ContactsList) {
					color.Printf("", "-----------------------------------\nContacts\n")
					color.Printf("d", fmt.Sprintf("\t\t[%d]\n%s\n", list.Total, list.UserIds))

					for _, contactUserId := range list.UserIds {
						client.User(contactUserId, func(user xingapi.User) {
							xingapi.PrintUser(user)
						})
					}
				})
			},
		},
	}

						// reader := bufio.NewReader(os.Stdin)
						// fmt.Print("Enter text: ")
						// text, _ := reader.ReadString('\n')
						// fmt.Println(text)


	app.Run(os.Args)
}

