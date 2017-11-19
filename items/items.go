package items

type Item struct {
	Corrupted         bool
	ExplicitMods      []string
	FlavorText        []string
	Height            int64  `json:"h"` // In stash spaces.
	Icon              string // URL.
	Identified        bool
	ItemLevel         int64 `json:"ilvl"`
	League            string
	LockedToCharacter bool // Chinese realm only.
	Name              string
	Note              string
	Properties        []Property
	Requirements      []Requirement
	Sockets           []Socket
	TalismanTier      int64
	TypeLine          string // Base type.
	Verified          bool
	Width             int64 `json:"w"` // In stash spaces.
	X                 int64 // Position in stash.
	Y                 int64 // Position in stash.

	// https://github.com/willroberts/loot/blob/master/forum/fixtures/items.json#L1238
	FrameType     int64       // Unknown.
	InventoryID   string      // Unknown.
	SocketedItems interface{} // See URL above.
	Support       bool        // Unknown.
}

type Property struct {
	DisplayMode int64 // Unknown.
	Name        string
	Values      [][]interface{} // FIXME.
}

type Requirement struct {
	DisplayMode int64 // Unknown.
	Name        string
	Values      [][]interface{} // FIXME.
}

type Socket struct {
	Attribute string `json:"attr"` // S, D, or I.
	Group     int64  // Numeric socket group (shows linked sockets).
}