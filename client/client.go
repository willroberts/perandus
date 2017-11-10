package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BaseURL string = "http://www.pathofexile.com/api/public-stash-tabs"
)

type Client interface {
	GetStashes(string) (*StashesResponse, error) // TODO: Make private.
	Poll() chan string
}

type client struct{}

func NewClient() Client {
	c := &client{}
	return c
}

func (c *client) Poll() chan string {
	ch := make(chan string)
	return ch
}

func (c *client) GetStashes(next string) (*StashesResponse, error) {
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
