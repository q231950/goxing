// xingapp.go

package main

import (
	"github.com/codegangsta/cli"
	"github.com/q231950/xingapi"
	"strconv"
)

// XINGApp encapsulates the client that connects to XING and holds repositories that provide data.
type XINGApp struct {
	cli.App
	client            xingapi.Client
	contactRepository xingapi.ContactRepository
	currentUser       xingapi.User
}

// NewApp creates a new app with default values
func NewApp(cliApp cli.App) *XINGApp {
	client := new(xingapi.XINGClient)
	contactRepository := xingapi.NewContactRepository(client)
	return &XINGApp{*cli.NewApp(), client, *contactRepository, nil}
}

func (app *XINGApp) LoadMeAction(c *cli.Context) {
	app.loadMe(func(me xingapi.User, err error) {
		if err == nil {
			app.currentUser = me
			xingapi.PrintUser(me)
		} else {
			xingapi.PrintError(err)
		}
	})
}

// LoadMeAction loads the logged in user
func (xa *XINGApp) loadMe(handler func(xingapi.User, error)) {
	if xa.currentUser != nil {
		handler(xa.currentUser, nil)
	} else {
		xa.client.Me(func(me xingapi.User, err error) {
			if err == nil {
				handler(me, nil)
			} else {
				handler(nil, err)
			}
		})
	}
}

// LoadContactsAction loads the contacts of the logged in user
func (app *XINGApp) LoadContactsAction(c *cli.Context) {
	app.loadMe(func(me xingapi.User, err error) {
		xingapi.Print(me.Id())
		userId := me.Id()
		app.contactRepository.Contacts(userId, func(users []*xingapi.User, err error) {
			if err != nil {
				xingapi.PrintError(err)
			} else {
				for index, user := range users {
					xingapi.PrintMessageWithParam("", strconv.Itoa(index+1)+".")
					xingapi.PrintUserOneLine(*user)
				}
			}
		})
	})
}

func (xa *XINGApp) LoadMessagesAction(c *cli.Context) {
	userId := c.Args().First()
	xa.client.Messages(userId, func(err error) {

	})
}
