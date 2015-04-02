// xingapp.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"xingapi"

	"github.com/codegangsta/cli"
	"github.com/str1ngs/ansi/color"
)

type XINGApp struct {
	cli.App
	s string
}

func NewApp(cliApp cli.App) *XINGApp {
	return &XINGApp{*cli.NewApp(), "some string"}
}

func (xa *XINGApp) loadMeAction(c *cli.Context) {

	client := new(xingapi.Client)
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
	client := new(xingapi.Client)

	// just to get the total
	client.ContactsList(userId, 0, 0, func(list xingapi.ContactsList, err error) {
		if err == nil {
			color.Printf("", fmt.Sprintf("-----------------------------------\n%d Contacts\n", list.Total))
			if 0 < list.Total {
				xa.requestLoadUsers(userId, list.Total, 0)
			}
		} else {
			xingapi.PrintError(err)
		}
	})
}

func (xa *XINGApp) LoadMessagesAction(c *cli.Context) {
	userId := c.Args().First()
	client := new(xingapi.Client)
	client.Messages(userId, func(err error) {

	})
}

func (xa *XINGApp) requestLoadUsers(request UsersRequest) {

	limit := request.Limit
	if request.Offset+request.Limit > request.Total {
		limit = request.Limit - (request.Offset + request.Limit - request.Total)
	}
	hint := ""
	if request.Offset == 0 {
		hint = "['y' or 'n']"
	}
	color.Printf("d", fmt.Sprintf("Load contacts (%d to %d)? %s\n", request.Offset, request.Offset+limit, hint))

	reader := bufio.NewReader(os.Stdin)
	xa.handleInputAndLoadContactsForUser(*reader, userId, limit, offset, total)
}

func (xa *XINGApp) handleInputAndLoadContactsForUser(reader bufio.Reader, userId string, limit int, offset int, total int) {
	text, _ := reader.ReadString('\n')
	if text == "y\n" {
		xa.loadUsers(userId, limit, offset, total)
	} else if text == "n\n" {
		// exit loop
	} else {
		println("Please enter 'y' or 'n'...")
		xa.requestLoadUsers(userId, total, offset)
	}
}

func (xa *XINGApp) loadUsers(userId string, limit int, offset int, total int) {
	client := new(xingapi.Client)
	client.ContactsList(userId, limit, offset, func(list xingapi.ContactsList, err error) {
		if err == nil {
			xa.loadAndPrintUsers(list)
			if offset+limit < total {
				xa.requestLoadUsers(userId, total, offset+len(list.UserIds))
			}
		} else {
			xingapi.PrintError(err)
		}
	})
}

func (xa *XINGApp) loadAndPrintUsers(list xingapi.ContactsList) {
	client := new(xingapi.Client)
	var waitGroup sync.WaitGroup
	for _, contactUserId := range list.UserIds {
		waitGroup.Add(1)
		go client.User(contactUserId, func(user xingapi.User, err error) {
			if err == nil {
				xingapi.PrintUserOneLine(user)
			} else {
				xingapi.PrintError(err)
			}
			defer waitGroup.Done()
		})
	}
	waitGroup.Wait()
}
