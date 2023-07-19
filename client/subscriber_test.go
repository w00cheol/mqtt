package client_test

import (
	"main/client"
	"testing"
)

func TestSubscribe(t *testing.T) {
	if err := client.Subscribe(); err != nil {
		t.Error(err)
	}
}
