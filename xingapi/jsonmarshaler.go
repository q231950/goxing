// jsonmarshaler.go
package xingapi

import (
	"io"
	"encoding/json"
	)

type UsersMarshaler interface {
	MarshalUsers(writer io.Writer, users Users) error
}

type UsersUnmarshaler interface {
	UnmarshalUsers(reader io.Reader) (Users, error)
}

type CredentialsMarshaler interface {
	MarshalCredentials(writer io.Writer, credentials Credentials) error	
}

type CredentialsUnmarshaler interface {
	UnmarshalCredentials(reader io.Reader) (Credentials, error)
}


type JSONMarshaler struct {}

func (JSONMarshaler) MarshalUsers(writer io.Writer, users Users) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(users)
}

func (JSONMarshaler) UnmarshalUsers(reader io.Reader) (Users, error) {
	decoder := json.NewDecoder(reader)
	var users Users
	err := decoder.Decode(&users)
	return users, err
}

func (JSONMarshaler) MarshalCredentials(writer io.Writer, credentials Credentials) error {
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(credentials)
	return err
}

func (JSONMarshaler) UnmarshalCredentials(reader io.Reader) (Credentials, error) {
	decoder := json.NewDecoder(reader)
	var credentials Credentials
	err := decoder.Decode(&credentials)
	return credentials, err
}