// user.go
package xingapi

import ()

type User struct {
	Name            string
	DisplayName     string `json:"display_name"`
	Id              string
	ActiveEmail     string    `json:"active_email"`
	Birthdate       Birthdate `json:"birth_date"`
	BusinessAddress Address   `json:"business_address"`
}

func (user User) String() string {
	return user.DisplayName + user.Name + " " + user.Id + " " + user.ActiveEmail + " - " + user.Birthdate.String() + user.BusinessAddress.String()
}
