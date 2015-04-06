// dummyclient.go

package xingapi

type DummyClient struct {
	Client
	Users []User
}

func (client *DummyClient) ContactsList(userID string, limit int, offset int, handler ContactsHandler) {
	list := new(ContactsList)
	list.UserIds = []string{"userId 1", "userId 2"}
	list.Total = 2
	handler(*list, nil)
}

func (client *DummyClient) User(contactUserId string, handler UserHandler) {
	handler(client.Users[0], nil)
}
