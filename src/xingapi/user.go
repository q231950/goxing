// user.go
package xingapi

import (
	"fmt"
)

type User interface {
	fmt.Stringer
	Name() string
	DisplayName() string
	Id() string
	ActiveEmail() string
	Birthdate() Birthdate
	BusinessAddress() Address
}

type XINGUser struct {
	name            string
	displayName     string `json:"display_name"`
	id              string
	activeEmail     string    `json:"active_email"`
	birthdate       Birthdate `json:"birth_date"`
	businessAddress Address   `json:"business_address"`
}

func (user XINGUser) String() string {
	return user.displayName + user.Name() + " " + user.Id() + " " + user.ActiveEmail() + " - " + user.Birthdate().String() + user.BusinessAddress().String()
}

func (user *XINGUser) Name() string {
	return user.name
}

func (user *XINGUser) DisplayName() string {
	return user.displayName
}

func (user *XINGUser) Id() string {
	return user.id
}

func (user *XINGUser) ActiveEmail() string {
	return user.activeEmail
}

func (user *XINGUser) Birthdate() Birthdate {
	return user.birthdate
}

func (user *XINGUser) BusinessAddress() Address {
	return user.businessAddress
}
