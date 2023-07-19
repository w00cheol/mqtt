package client_test

import (
	"main/client"
	"testing"
)

func TestPublish(t *testing.T) {
	if err := client.Publish(); err != nil {
		t.Error(err)
	}
}
