// xingapp.go
package main

import (
	"github.com/codegangsta/cli"
	"strconv"
	"xingapi"
)

type XINGApp struct {
	cli.App
	s                 string
	client            xingapi.Client
	contactRepository xingapi.ContactRepository
}

func NewApp(cliApp cli.App) *XINGApp {
	client := new(xingapi.XINGClient)
	contactRepository := xingapi.NewContactRepository(client)
	return &XINGApp{*cli.NewApp(), "some string", client, *contactRepository}
}

func (xa *XINGApp) loadMeAction(c *cli.Context) {
	xa.client.Me(func(me xingapi.User, err error) {
		if err == nil {
			xingapi.PrintUser(me)
		} else {
			xingapi.PrintError(err)
		}
	})
}

func (xa *XINGApp) LoadContactsAction(c *cli.Context) {
	userId := c.Args().First()
	xa.contactRepository.Contacts(userId, func(users []*xingapi.User, err error) {
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
	xa.client.Messages(userId, func(err error) {

	})
}
