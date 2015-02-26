// client.go
package xingapi

import (
	"fmt"
	"net/url"
	"io"
	"io/ioutil"
	"log"
	)

type MeHandler func(me User)

type Client struct {
	OAuthConsumer OAuthConsumer
	meHandler MeHandler
}

func (client *Client)Me(handler MeHandler) {
	
	client.meHandler = handler
	consumer := new(OAuthConsumer)
	client.OAuthConsumer = *consumer
	// client.OAuthConsumer.Connect()
	client.OAuthConsumer.Get("/v1/users/me", url.Values{}, client.MeResponseHandler)
}

func (client *Client)Contacts(userID string, handler func()) {	
	client.OAuthConsumer.Get("/v1/users/"+ userID + "/contacts", url.Values{}, client.ContactsResponseHandler)
}

func (client *Client)ContactsResponseHandler(reader io.Reader) {
	robots, err := ioutil.ReadAll(reader)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Printf("%s", robots)
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
