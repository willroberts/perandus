package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client interface {
	getStashes(string) (*StashesResponse, error)
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

func (c *client) getStashes(next string) (*StashesResponse, error) {
	url := BaseURL
	if next != "" {
		url = fmt.Sprintf("%s?id=%s", BaseURL, next)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var s StashesResponse
	err = json.Unmarshal(b, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
