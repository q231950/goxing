// xingapp.go
package main

import (
	"xingapi"

	"github.com/codegangsta/cli"
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
	repo.Contacts(userId, func(list []xingapi.User, err error) {

	})
}

func (xa *XINGApp) LoadMessagesAction(c *cli.Context) {
	userId := c.Args().First()
	client := new(xingapi.XINGClient)
	client.Messages(userId, func(err error) {

	})
}
