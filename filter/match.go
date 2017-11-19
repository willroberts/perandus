package filter

import (
	"regexp"

	"github.com/willroberts/perandus/items"
)

var (
	setRemover = regexp.MustCompile("<<.*>>")
)

func (f *filter) Matches(i items.Item) bool {
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

func matchesLeague(i items.Item, league string) bool {
	return i.League == league
}

func matchesName(i items.Item, name string) bool {
	return pruneItemName(i.Name) == name
}

func matchesMinPrice(i items.Item, price string) bool {
	return true
}

func matchesMaxPrice(i items.Item, price string) bool {
	return true
}

func pruneItemName(name string) string {
	b := setRemover.ReplaceAll([]byte(name), []byte(""))
	return string(b)
}
