package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/willroberts/perandus/models"
)

const (
	baseURL string = "http://www.pathofexile.com/api/public-stash-tabs"
)

// Client is a rate-limited HTTP client for PoE's API.
type Client interface {
	Poll() (chan models.Item, chan error)
}

type client struct {
	NextChangeID string
	ItemHistory  map[string]struct{}
}

// New initializes and returns a Client.
func New(nextChangeID string) Client {
	return &client{
		NextChangeID: nextChangeID,
		ItemHistory:  make(map[string]struct{}),
	}
}

// StashesResponse reprensents a JSON response from the PoE API.
type StashesResponse struct {
	NextChangeID string `json:"next_change_id"`
	Stashes      []models.Stash
}

func (c *client) getOne(changeID string) (*StashesResponse, error) {
	url := fmt.Sprintf("%s?id=%s", baseURL, changeID)
	resp, err := http.Get(url)
	if err != nil {
		return &StashesResponse{}, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &StashesResponse{}, err
	}

	var s StashesResponse
	if err := json.Unmarshal(b, &s); err != nil {
		return &StashesResponse{}, err
	}

	return &s, nil
}

func (c *client) addToHistory(i models.Item) {
	c.ItemHistory[i.ID] = struct{}{}
}

func (c *client) isInHistory(i models.Item) bool {
	if _, ok := c.ItemHistory[i.ID]; ok {
		return true
	}
	return false
}
