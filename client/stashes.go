package client

import "github.com/willroberts/perandus/items"

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
