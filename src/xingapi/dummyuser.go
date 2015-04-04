// dummyuser.go

package xingapi

type DummyUser struct {
	User
}

func (user DummyUser) String() string {
	return "dummy user"
}
