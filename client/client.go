package client

type Client interface {
	getOne(string) (*StashesResponse, error)
	Poll()
}

type client struct {
	NextChangeID string
}

func New(nextChangeID string) Client {
	return &client{
		NextChangeID: nextChangeID,
	}
}
