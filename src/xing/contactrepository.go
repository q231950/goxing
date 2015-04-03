// contactrespository.go

package main

import (
	"xingapi"
)

type ContactRepository struct {
}

func (repo *ContactRepository) Contacts(contacts func(list []xingapi.User, err error)) {

}
