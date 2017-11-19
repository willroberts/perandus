package client

import (
	"log"

	"github.com/willroberts/perandus/filter"
)

type Client interface {
	getOne(string) (*StashesResponse, error)
	Poll()
}

type client struct {
	NextChangeID string
	Filter       filter.Filter
	FilterQueue  chan Stash
}

func New(nextChangeID string) Client {
	c := &client{
		NextChangeID: nextChangeID,
		FilterQueue:  make(chan Stash),
	}
	f, err := filter.New()
	if err != nil {
		log.Fatal("exit from client.New")
	}
	c.Filter = f
	go c.runFilterWorker()
	return c
}
