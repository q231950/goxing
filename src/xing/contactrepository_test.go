// contactsrepository_test.go

package main

import (
	"strconv"
	"sync"
	"testing"
	"xingapi"
)

func TestGetContacts(t *testing.T) {
	client := new(xingapi.DummyClient)
	dummyUsers := make([]xingapi.User, 2)
	dummyUsers[0] = new(xingapi.DummyUser)
	dummyUsers[1] = new(xingapi.DummyUser)
	client.Users = dummyUsers
	repository := NewContactRepository(client)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	repository.Contacts("some user id", func(list []xingapi.User, err error) {
		if len(list) != 2 {
			t.Error("Expected '2' but got '" + strconv.Itoa(len(list)) + "'")
		}
		waitGroup.Done()
	})
	waitGroup.Wait()
}
