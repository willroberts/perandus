package models

import "regexp"

// Item models a generic item in PoE.
type Item struct {
	Corrupted         bool
	ExplicitMods      []string
	FlavorText        []string `json:"flavourText"`
	FrameType         int64
	Height            int64 `json:"h"` // In stash spaces.
	ID                string
	Icon              string // URL.
	Identified        bool
	ImplicitMods      []string
	InventoryID       string
	ItemLevel         int64 `json:"ilvl"`
	League            string
	LockedToCharacter bool // Chinese realm only.
	Name              string
	Note              string
	Properties        []Property
	Requirements      []Requirement
	SocketedItems     interface{}
	Sockets           []Socket
	Support           bool
	TalismanTier      int64
	TypeLine          string // Base type.
	Verified          bool
	Width             int64 `json:"w"` // In stash spaces.
	X                 int64 // Position in stash.
	Y                 int64 // Position in stash.
}

// Property stores a pair of name and values.
type Property struct {
	DisplayMode int64 // Unknown.
	Name        string
	Values      [][]interface{} // FIXME.
}

// Requirement stores attribute requirements.
type Requirement struct {
	DisplayMode int64 // Unknown.
	Name        string
	Values      [][]interface{} // FIXME.
}

// Socket stores information about sockets and links.
type Socket struct {
	Attribute string `json:"attr"` // S, D, or I.
	Group     int64  // Numeric socket group (shows linked sockets).
}

var tagFinder = regexp.MustCompile("<<.*>>")

// StripLocalizationTags removes localization data from item names.
func StripLocalizationTags(name string) string {
	b := tagFinder.ReplaceAll([]byte(name), []byte(""))
	return string(b)
}
