package filter

import (
	toml "github.com/pelletier/go-toml"
	"github.com/willroberts/perandus/items"
)

// Filter reads settings and compares items to the configured parameters.
type Filter interface {
	Matches(items.Item) bool
}

type filter struct {
	Settings *toml.Tree
	League   string
	ItemName string
	MinPrice string
	MaxPrice string
}

// New initializes and returns a Filter.
func New() (Filter, error) {
	f := &filter{}
	s, err := f.parseSettings()
	if err != nil {
		return f, err
	}
	f.Settings = s
	return f, nil
}
