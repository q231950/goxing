// client_test.go

package xingapi

import (
	"testing"
)

func TestClient(t *testing.T) {
	c := new(Client)
	expectedName := "XING API client"
	if c.Name() != expectedName {
		t.Error("Expected '" + expectedName + "' but got '" + c.Name() + "'")
	}
}
