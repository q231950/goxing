// contactsrepository_test.go

package main

import (
	"testing"
	"xingapi"
)

func TestGetContacts(t *testing.T) {
	repository := new(ContactRepository)

	repository.Contacts(func(list []xingapi.User, err error) {

	})
	//	if c.Name() != expectedName {
	//		t.Error("Expected '" + expectedName + "' but got '" + c.Name() + "'")
	//	}
}
