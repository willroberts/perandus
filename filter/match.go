package filter

import (
	"github.com/willroberts/perandus/models"
)

func (f *filter) matches(i models.Item) bool {
	if i.League != f.league {
		return false
	}

	if i.Name != f.itemName {
		return false
	}

	return true
}
