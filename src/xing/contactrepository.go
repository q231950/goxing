// contactrespository.go

package main

import (
	"bufio"
	"fmt"
	"github.com/str1ngs/ansi/color"
	"os"
	"sync"
	"xingapi"
)

type ContactRepository struct {
	client xingapi.Client
}

func NewContactRepository(client xingapi.Client) *ContactRepository {
	return &ContactRepository{client}
}

func (repo *ContactRepository) Contacts(userId string, contacts func(list []xingapi.User, err error)) {
	repo.client.ContactsList(userId, 0, 0, func(list xingapi.ContactsList, err error) {
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
				repo.requestLoadUsers(request)
			}
		} else {
			xingapi.PrintError(err)
		}
	})
}

func (repo *ContactRepository) requestLoadUsers(request xingapi.UsersRequest) {

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
	repo.handleInputAndLoadContactsForUser(*reader, newRequest)
}

func (repo *ContactRepository) handleInputAndLoadContactsForUser(reader bufio.Reader, request xingapi.UsersRequest) {
	text, _ := reader.ReadString('\n')
	if text == "y\n" {
		repo.loadUsers(request)
	} else if text == "n\n" {
		// exit loop
		request.Completion(nil)
	} else {
		println("Please enter 'y' or 'n'...")
		repo.requestLoadUsers(request)
	}
}

func (repo *ContactRepository) loadUsers(request xingapi.UsersRequest) {
	repo.client.ContactsList(request.UserId, request.Limit, request.Offset, func(list xingapi.ContactsList, err error) {
		if err == nil {
			repo.loadAndPrintUsers(list)
			if !request.IsFinal() {
				newRequest := xingapi.UsersRequest{request.UserId, request.Limit, request.Offset + len(list.UserIds), request.Total, request.Completion}
				repo.requestLoadUsers(newRequest)
			} else {
				// finished final request without errors
				request.Completion(nil)
			}
		} else {
			request.Completion(err)
		}
	})
}

func (repo *ContactRepository) loadAndPrintUsers(list xingapi.ContactsList) {
	var waitGroup sync.WaitGroup
	for _, contactUserId := range list.UserIds {
		waitGroup.Add(1)
		go repo.client.User(contactUserId, func(user xingapi.User, err error) {
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
