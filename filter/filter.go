package filter

import (
	"github.com/willroberts/perandus/items"
)

// Filter reads settings and compares items to the configured parameters.
type Filter interface {
	Matches(items.Item) bool
}

type filter struct {
	League   string
	ItemName string
	MinPrice string
	MaxPrice string
}

// New initializes and returns a Filter.
func New() (Filter, error) {
	f := &filter{}
	if err := f.parseSettings(); err != nil {
		return f, err
	}
	return f, nil
}
