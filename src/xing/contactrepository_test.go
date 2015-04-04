// contactsrepository_test.go

package main

import (
	"testing"
	"xingapi"
)

func TestGetContacts(t *testing.T) {
	client := new(xingapi.DummyClient)
	dummyUsers := make([]xingapi.User, 2)
	dummyUsers[0] = new(xingapi.DummyUser)
	dummyUsers[1] = new(xingapi.DummyUser)

	repository := NewContactRepository(client)

	repository.Contacts(func(list []xingapi.User, err error) {

	})
	//	if c.Name() != expectedName {
	//		t.Error("Expected '" + expectedName + "' but got '" + c.Name() + "'")
	//	}
}
