package client

type Client interface {
	getOne(string) (*StashesResponse, error)
	Poll()
}

type client struct {
	NextChangeID string
	FilterQueue  chan Stash
}

func New(nextChangeID string) Client {
	c := &client{
		NextChangeID: nextChangeID,
		FilterQueue:  make(chan Stash),
	}
	go c.runFilterWorker()
	return c
}
