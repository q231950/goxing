// xingapp.go
package main

import (
	"github.com/codegangsta/cli"
	"github.com/str1ngs/ansi/color"
	"xingapi"
	"fmt"
	"bufio"
	"os"
	)

type XINGApp struct {
	cli.App
}

func init() {
	println("init XINGApp")
}

func (xa *XINGApp) loadMeAction(c *cli.Context) {

	client := new(xingapi.Client)	
	client.Me(func(me xingapi.User) {
		xingapi.PrintUser(me)
	})
}

func (xa *XINGApp) loadUserAction(c *cli.Context) {
	userId := c.Args().First()
	client := new(xingapi.Client)
	client.ContactsList(userId, func(list xingapi.ContactsList) {
		color.Printf("", "-----------------------------------\nContacts\n")
		
		color.Printf("d", fmt.Sprintf("Load %d users?\n", list.Total))
		xa.loadUsers(list)
	})
}

func (xa *XINGApp) loadUsers(list xingapi.ContactsList) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if text == "y\n" {
		client := new(xingapi.Client)
		for _, contactUserId := range list.UserIds {
			client.User(contactUserId, func(user xingapi.User) {
				xingapi.PrintUser(user)
			})
		}
	} else if text == "n\n" {
		// exit loop
	} else {
		println("Please enter 'y' or 'n'.")
		xa.loadUsers(list)
	}
}

