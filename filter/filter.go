package filter

import (
	toml "github.com/pelletier/go-toml"
	"github.com/willroberts/perandus/items"
)

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

func New() (Filter, error) {
	f := &filter{}
	s, err := f.parseSettings()
	if err != nil {
		return f, err
	}
	f.Settings = s
	return f, nil
}
