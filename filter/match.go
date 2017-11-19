package filter

import (
	"github.com/willroberts/perandus/models"
)

func (f *filter) Matches(i models.Item) bool {
	if !matchesLeague(i, f.League) {
		return false
	}
	if !matchesName(i, f.ItemName) {
		return false
	}
	if !matchesMinPrice(i, f.MinPrice) {
		return false
	}
	if !matchesMaxPrice(i, f.MaxPrice) {
		return false
	}
	return true
}

func matchesLeague(i models.Item, league string) bool {
	return i.League == league
}

func matchesName(i models.Item, name string) bool {
	return i.Name == name
}

func matchesMinPrice(i models.Item, price string) bool {
	return true
}

func matchesMaxPrice(i models.Item, price string) bool {
	return true
}
