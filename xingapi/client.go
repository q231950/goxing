// client.go
package xingapi

import (
	"fmt"
	"net/url"
	"io"
	"io/ioutil"
	"log"
	)

type UserHandler func(user User)
type ContactsHandler func(list ContactsList)

type Client struct {
	OAuthConsumer OAuthConsumer
	meHandler UserHandler
	contactsHandler ContactsHandler
}

func (client *Client)Me(handler UserHandler) {
	client.meHandler = handler
	consumer := new(OAuthConsumer)
	client.OAuthConsumer = *consumer
	client.OAuthConsumer.Get("/v1/users/me", url.Values{}, client.MeResponseHandler)
}

func (client *Client)MeResponseHandler(reader io.Reader) {
	users, err := client.readUsers(reader)
	if err == nil {
		me := users.Users[0]
		client.meHandler(*me)
	}
}

func (client *Client)readUsers(reader io.Reader) (Users, error) {
	var unmarshaler UsersUnmarshaler
	unmarshaler = JSONMarshaler{}
	return unmarshaler.UnmarshalUsers(reader)
}

func (client *Client)ContactsList(userID string, handler ContactsHandler) {	
	client.contactsHandler = handler
	consumer := new(OAuthConsumer)
	client.OAuthConsumer = *consumer
	client.OAuthConsumer.Get("/v1/users/"+ userID + "/contacts", url.Values{}, client.ContactsResponseHandler)
}

func (client *Client)ContactsResponseHandler(reader io.Reader) {
	list, err := client.readContactsList(reader)
	if err == nil {
		client.contactsHandler(list)
	} else {
		println(err.Error())
	}
}

func (client *Client)readContactsList(reader io.Reader) (list ContactsList, err error) {
	var unmarshaler ContactsListUnmarshaler
	unmarshaler = JSONMarshaler{}
	return unmarshaler.UnmarshalContactsList(reader)
}

func (client *Client) User(id string, handler UserHandler) {
	consumer := new(OAuthConsumer)
	client.OAuthConsumer = *consumer
	client.OAuthConsumer.Get("/v1/users/" + id, url.Values{}, func(reader io.Reader){
		// robots, err := ioutil.ReadAll(reader)
		// if err != nil {
		//     log.Fatal(err)
		// }
		// fmt.Printf("%s", robots)
		var unmarshaler UserUnmarshaler
		unmarshaler = JSONMarshaler{}
		user, _ := unmarshaler.UnmarshalUser(reader)
		handler(user)
	})
}

func (client *Client)Bla(reader io.Reader) {
	robots, err := ioutil.ReadAll(reader)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}