// xingapp.go
package main

import (
	"github.com/codegangsta/cli"
	"strconv"
	"xingapi"
)

type XINGApp struct {
	cli.App
	s string
}

func NewApp(cliApp cli.App) *XINGApp {
	return &XINGApp{*cli.NewApp(), "some string"}
}

func (xa *XINGApp) loadMeAction(c *cli.Context) {
	client := new(xingapi.XINGClient)
	client.Me(func(me xingapi.User, err error) {
		if err == nil {
			xingapi.PrintUser(me)
		} else {
			xingapi.PrintError(err)
		}
	})
}

func (xa *XINGApp) LoadContactsAction(c *cli.Context) {
	userId := c.Args().First()
	client := new(xingapi.XINGClient)
	repo := NewContactRepository(client)
	repo.Contacts(userId, func(users []*xingapi.User, err error) {
		if err != nil {
			xingapi.PrintError(err)
		} else {
			for index, user := range users {
				xingapi.PrintMessageWithParam("", strconv.Itoa(index+1)+".")
				xingapi.PrintUserOneLine(*user)
			}
		}
	})
}

func (xa *XINGApp) LoadMessagesAction(c *cli.Context) {
	userId := c.Args().First()
	client := new(xingapi.XINGClient)
	client.Messages(userId, func(err error) {

	})
}
