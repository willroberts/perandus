package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/willroberts/loot/items"
)

const (
	BaseURL string = "http://www.pathofexile.com/api/public-stash-tabs"
)

type StashesResponse struct {
	NextChangeID string `json:"next_change_id"`
	Stashes      []Stash
}

type Stash struct {
	AccountName       string
	LastCharacterName string
	ID                string
	Stash             string
	Items             []items.Item
	Public            bool
}

type Client interface {
	GetStashes(string) (*StashesResponse, error) // TODO: Make private.
	Poll() chan string
}

type client struct{}

func NewClient() Client {
	c := &client{}
	return c
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
