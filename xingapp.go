package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/codegangsta/cli"
	"github.com/q231950/xingapi"
	"github.com/str1ngs/ansi/color"
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

// LoadMeAction is the action to fetch the signed in user
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

func (app *XINGApp) loadMe(handler func(xingapi.User, error)) {
	if app.currentUser != nil {
		handler(app.currentUser, nil)
	} else {
		app.client.Me(func(me xingapi.User, err error) {
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
		app.contactRepository.Contacts(me.ID(), func(users []*xingapi.User, err error) {
			color.Printf("", fmt.Sprintf("-----------------------------------\n%d Contacts\n", len(users)))
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

// LoadContactAction loads the contact with the given name
func (app *XINGApp) LoadContactAction(c *cli.Context) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                   // Start the spinner

	app.loadMe(func(me xingapi.User, err error) {
		contactName := c.Args().First()
		app.contactRepository.Contact(contactName, func() {
			s.Stop()
		})
	})
}

// LoadMessagesAction is the action to load the messages of the signed in user
func (app *XINGApp) LoadMessagesAction(c *cli.Context) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                   // Start the spinner

	userID := c.Args().First()
	app.client.Messages(userID, func(err error) {
		s.Stop()
	})
}
