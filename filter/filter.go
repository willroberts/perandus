package filter

import (
	"github.com/willroberts/perandus/models"
)

// Filter reads settings and compares items to the configured parameters.
type Filter interface {
	FilterItems(chan models.Item) chan models.Item
}

type filter struct {
	league   string
	itemName string
}

// New initializes and returns a Filter.
func New() (Filter, error) {
	f := &filter{}
	if err := f.parseSettings(); err != nil {
		return f, err
	}
	return f, nil
}

func (f *filter) FilterItems(in chan models.Item) chan models.Item {
	out := make(chan models.Item)

	go func() {
		for {
			select {
			case i := <-in:
				if f.matches(i) && i.Note != "" {
					out <- i
				}
			}
		}
	}()

	return out
}
