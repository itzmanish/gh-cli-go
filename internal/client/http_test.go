package client

import (
	"testing"
	"time"

	"github.com/itzmanish/gh-cli-go/internal"
)

func TestNewClient(t *testing.T) {
	timeout := time.Second * 5
	c := newClient(timeout)
	if c == nil {
		t.Error("Expected client but got nil")
	}
	if c.Timeout != timeout {
		t.Errorf("Client should intialize with given timeout: %s but got initialized with %s", timeout, c.Timeout)
	}
}

func TestNewRequest(t *testing.T) {
	err := internal.LoadConfig("../../samples")
	if err != nil {
		t.Error(err)
	}
	res, err := NewRequestWithAuthentication(CurrentUserURL)
	if err != nil {
		t.Error(err)
	}
	if res == nil {
		t.Error("Expected some response but got nil")
	}
}
