// dummyclient.go

package xingapi

type DummyClient struct {
	Client
	Users []User
}

func (client *DummyClient) ContactsList(userID string, limit int, offset int, handler ContactsHandler) {
	list := new(ContactsList)
	handler(*list, nil)
}
