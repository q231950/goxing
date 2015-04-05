// contactrespository.go

package main

import (
	"fmt"
	"github.com/str1ngs/ansi/color"

	"sync"
	"xingapi"
)

type ContactRepository struct {
	client xingapi.Client
}

func NewContactRepository(client xingapi.Client) *ContactRepository {
	return &ContactRepository{client}
}

func (repo *ContactRepository) Contacts(userId string, contactsHandler func(list []*xingapi.User, err error)) {
	repo.client.ContactsList(userId, 0, 0, func(list xingapi.ContactsList, err error) {
		if err == nil {
			color.Printf("", fmt.Sprintf("-----------------------------------\n%d Contacts\n", list.Total))
			if 0 < list.Total {
				limit := 50
				request := xingapi.UsersRequest{userId, limit, 0, list.Total, contactsHandler}
				repo.requestLoadUsers(request)
			} else {
				contactsHandler([]*xingapi.User{}, nil)
			}
		} else {
			contactsHandler(nil, err)
		}
	})
}

func (repo *ContactRepository) requestLoadUsers(request xingapi.UsersRequest) {

	limit := request.Limit
	if request.Offset+request.Limit > request.Total {
		limit = request.Limit - (request.Offset + request.Limit - request.Total)
	}

	newRequest := xingapi.UsersRequest{request.UserId,
		limit,
		request.Offset,
		request.Total,
		request.Completion}
	repo.loadUsers([]*xingapi.User{}, newRequest)
}

func (repo *ContactRepository) loadUsers(users []*xingapi.User, request xingapi.UsersRequest) {
	repo.client.ContactsList(request.UserId, request.Limit, request.Offset, func(list xingapi.ContactsList, err error) {
		if err == nil {
			repo.loadUserDetails(list, func(loadedUsers []*xingapi.User, err error) {
				users = append(users, loadedUsers...)
				if !request.IsFinal() {
					newRequest := xingapi.UsersRequest{request.UserId, request.Limit, request.Offset + len(list.UserIds), request.Total, request.Completion}
					repo.loadUsers(users, newRequest)
				} else {
					// finished final request without errors
					request.Completion(users, nil)
				}
			})
		} else {
			request.Completion(nil, err)
		}
	})
}

func (repo *ContactRepository) loadUserDetails(list xingapi.ContactsList, loadedUsers func(userList []*xingapi.User, err error)) {
	users := []*xingapi.User{}
	//	var err error
	var waitGroup sync.WaitGroup
	for _, contactUserId := range list.UserIds {
		waitGroup.Add(1)
		go repo.client.User(contactUserId, func(user xingapi.User, cerr error) {
			if cerr == nil {
				users = append(users, &user)
			} else {
				xingapi.PrintError(cerr)
				//				err = cerr
			}
			defer waitGroup.Done()
		})
	}
	waitGroup.Wait()
	loadedUsers(users, nil)
}
