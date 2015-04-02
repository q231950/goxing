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
				limit := 20
				request := xingapi.UsersRequest{userId, limit, 0, list.Total, func(err error) {
					if err != nil {
						xingapi.PrintError(err)
					} else {
						println("done")
					}
				}}
				xa.requestLoadUsers(request)
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

func (xa *XINGApp) requestLoadUsers(request xingapi.UsersRequest) {

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
	newRequest := xingapi.UsersRequest{request.UserId, limit, request.Offset, request.Total, request.Completion}
	xa.handleInputAndLoadContactsForUser(*reader, newRequest)
}

func (xa *XINGApp) handleInputAndLoadContactsForUser(reader bufio.Reader, request xingapi.UsersRequest) {
	text, _ := reader.ReadString('\n')
	if text == "y\n" {
		xa.loadUsers(request)
	} else if text == "n\n" {
		// exit loop
		request.Completion(nil)
	} else {
		println("Please enter 'y' or 'n'...")
		xa.requestLoadUsers(request)
	}
}

func (xa *XINGApp) loadUsers(request xingapi.UsersRequest) {
	client := new(xingapi.Client)
	client.ContactsList(request.UserId, request.Limit, request.Offset, func(list xingapi.ContactsList, err error) {
		if err == nil {
			xa.loadAndPrintUsers(list)
			if !request.IsFinal() {
				newRequest := xingapi.UsersRequest{request.UserId, request.Limit, request.Offset + len(list.UserIds), request.Total, request.Completion}
				xa.requestLoadUsers(newRequest)
			} else {
				// finished final request without errors
				request.Completion(nil)
			}
		} else {
			request.Completion(err)
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
