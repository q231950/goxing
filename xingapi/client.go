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

type UserHandler func(User)
type ContactsHandler func(ContactsList, error)

type Client struct {
	OAuthConsumer OAuthConsumer
}

func (client *Client)Me(handler UserHandler) {
	var me User
	consumer := new(OAuthConsumer)
	client.OAuthConsumer = *consumer
	client.OAuthConsumer.Get("/v1/users/me", url.Values{}, func(reader io.Reader) {
		var unmarshaler UsersUnmarshaler
		unmarshaler = JSONMarshaler{}

		users, err := unmarshaler.UnmarshalUsers(reader)
		if err == nil {
			me = *users.Users[0]
			handler(me)
		}
	})
}

func (client *Client) ContactsList(userID string, limit int, offset int, handler ContactsHandler) {	
	consumer := new(OAuthConsumer)
	client.OAuthConsumer = *consumer
	v := url.Values{}
	v.Set("limit", strconv.Itoa(limit))
	v.Set("offset", strconv.Itoa(offset))
	v.Set("order_by", "last_name")
	client.OAuthConsumer.Get("/v1/users/"+ userID + "/contacts", v, func(reader io.Reader) {

		var unmarshaler ContactsListUnmarshaler
		unmarshaler = JSONMarshaler{}
		list, err := unmarshaler.UnmarshalContactsList(reader)
		handler(list, err)
	})
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
