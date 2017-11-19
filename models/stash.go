package models

// Stash models the contents of a single public stash tab.
type Stash struct {
	AccountName       string
	LastCharacterName string
	ID                string
	Stash             string
	Items             []Item
	Public            bool
}
