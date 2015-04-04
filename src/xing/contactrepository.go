// contactrespository.go

package main

import (
	"xingapi"
)

type ContactRepository struct {
	client xingapi.Client
}

func NewContactRepository(client xingapi.Client) *ContactRepository {
	return &ContactRepository{client}
}

func (repo *ContactRepository) Contacts(contacts func(list []xingapi.User, err error)) {

}
