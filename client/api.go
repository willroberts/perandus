package client

import (
	"github.com/willroberts/perandus/items"
)

const (
	baseURL string = "http://www.pathofexile.com/api/public-stash-tabs"
)

// StashesResponse models the response from the PoE API.
type StashesResponse struct {
	NextChangeID string `json:"next_change_id"`
	Stashes      []Stash
}

// Stash models the contents of a single public stash tab.
type Stash struct {
	AccountName       string
	LastCharacterName string
	ID                string
	Stash             string
	Items             []items.Item
	Public            bool
}
