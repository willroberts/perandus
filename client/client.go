package client

import (
	"github.com/willroberts/perandus/filter"
	"github.com/willroberts/perandus/models"
)

// Client is a rate-limited HTTP client for PoE's API.
type Client interface {
	getOne(string) (*StashesResponse, error)
	Poll() error
}

type client struct {
	NextChangeID string
	ItemHistory  map[string]struct{}
	Filter       filter.Filter
	FilterQueue  chan models.Stash
}

// New initializes and returns a Client.
func New(nextChangeID string) (Client, error) {
	c := &client{
		NextChangeID: nextChangeID,
		ItemHistory:  make(map[string]struct{}),
		FilterQueue:  make(chan models.Stash),
	}
	f, err := filter.New()
	if err != nil {
		return c, err
	}
	c.Filter = f
	go c.runFilterWorker()
	return c, nil
}
