package main

import (
	"github.com/codegangsta/cli"
	"github.com/str1ngs/ansi/color"
	"os"
	"xingapi"
	"fmt"
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
						color.Printf("+g", fmt.Sprintf("-----------------------------------\n%s:\n", me.DisplayName))
						color.Printf("g", fmt.Sprintf("Email address:\t\t%s\nDate of birth:\t\t%s\n", me.ActiveEmail, me.Birthdate))

						// reader := bufio.NewReader(os.Stdin)
						// fmt.Print("Enter text: ")
						// text, _ := reader.ReadString('\n')
						// fmt.Println(text)
					})
			},
		},
		{
			Name:      "complete",
			ShortName: "c",
			Usage:     "complete a task on the list",
			Action: func(c *cli.Context) {
				println("completed task: ", c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}

