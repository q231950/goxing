// client.go
package xingapi

import (
	"fmt"
	"net/url"
	"io"
	"io/ioutil"
	"log"
	"strconv"
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

func (client *Client) ContactsList(userID string, limit int, offset int, handler ContactsHandler) {	
	client.contactsHandler = handler
	consumer := new(OAuthConsumer)
	client.OAuthConsumer = *consumer
	v := url.Values{}
	v.Set("limit", strconv.Itoa(limit))
	v.Set("offset", strconv.Itoa(offset))
	v.Set("order_by", "last_name")
	client.OAuthConsumer.Get("/v1/users/"+ userID + "/contacts", v, client.ContactsResponseHandler)
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
		var unmarshaler UserUnmarshaler
		unmarshaler = JSONMarshaler{}
		user, _ := unmarshaler.UnmarshalUser(reader)
		handler(user)
	})
}

// GET /v1/users/:user_id/conversations
func (client *Client)Messages(userId string, handler func(err error)) {
	client.OAuthConsumer.Get("/v1/users/" + userId + "/conversations", url.Values{}, func(reader io.Reader){
		robots, readError := ioutil.ReadAll(reader)
		
		println(fmt.Sprintf("%s", robots))

		handler(readError)
	})
}

func (client *Client)Bla(reader io.Reader) {
	robots, err := ioutil.ReadAll(reader)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}