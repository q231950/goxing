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

func (xa *XINGApp) LoadContactsAction(c *cli.Context) {
	userId := c.Args().First()
	client := new(xingapi.Client)
	
	// just to get the total
	client.ContactsList(userId, 0, 0, func(list xingapi.ContactsList) {
		color.Printf("", fmt.Sprintf("-----------------------------------\n%d Contacts\n", list.Total))
		if 0 < list.Total {
			xa.requestLoadUsers(userId, list.Total, 0)
		}
	})
}

func (xa *XINGApp)LoadMessagesAction(c *cli.Context) {
	userId := c.Args().First()
	client := new(xingapi.Client)
	client.Messages(userId, func(err error) {
		
	})
}

func (xa *XINGApp) requestLoadUsers(userId string, total int, offset int) {

	limit := 20
	if offset + limit > total {
		limit = limit - (offset + limit - total)
	}
	hint := ""
	if offset == 0 {
		hint = "['y' or 'n']"
	}
	color.Printf("d", fmt.Sprintf("Load users (%d to %d)? %s\n", offset, offset + limit, hint))

	client := new(xingapi.Client)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if text == "y\n" {
		client.ContactsList(userId, limit, offset, func(list xingapi.ContactsList) {
			xa.loadUsers(list)
			if offset+limit < total {
				xa.requestLoadUsers(userId, total, offset + len(list.UserIds))
			}
		})
	} else if text == "n\n" {
		// exit loop
	} else {
		println("Please enter 'y' or 'n'...")
		xa.requestLoadUsers(userId, total, offset)
	}
}

func (xa *XINGApp) loadUsers(list xingapi.ContactsList) {
	
	client := new(xingapi.Client)
	for _, contactUserId := range list.UserIds {
		client.User(contactUserId, func(user xingapi.User) {
			xingapi.PrintUserOneLine(user)
		})
	}
}

