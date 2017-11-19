package client

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/willroberts/perandus/filter"
	"github.com/willroberts/perandus/items"
)

func (c *client) runFilterWorker() {
	for {
		s := <-c.FilterQueue

		if len(s.Items) == 0 {
			continue
		}

		for _, i := range s.Items {
			// Strip localization tags from item names before proceeding.
			// It may be better to do this elsewhere.
			i.Name = filter.StripLocalizationTags(i.Name)

			// Compute the hash and store it alongside the item.
			b, err := json.Marshal(i)
			if err != nil {
				log.Println("Error hashing item:", err.Error())
				continue
			}
			hashBytes := md5.Sum(b)
			i.Hash = hex.EncodeToString(hashBytes[:])

			// Keep track of this item in the client history.
			if !c.IsInHistory(i) {
				c.AddToHistory(i)
			} else {
				// We've already seen this item; don't alert.
				continue
			}

			if c.Filter.Matches(i) && i.Note != "" {
				log.Println("ALERT! Found matching item:")
				log.Println("\tName:", i.Name)
				log.Println("\tPrice:", i.Note)
				log.Println("\tSeller:", s.LastCharacterName)
				log.Println("\tHash:", i.Hash)
			}
		}
	}
}

func (c *client) AddToHistory(i items.Item) {
	c.ItemHistory[i.Hash] = struct{}{}
}

func (c *client) IsInHistory(i items.Item) bool {
	if _, ok := c.ItemHistory[i.Hash]; ok {
		return true
	}
	return false
}
