package client

import (
	"log"

	"github.com/willroberts/perandus/models"
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
			i.Name = models.StripLocalizationTags(i.Name)

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
				log.Println("\tID:", i.ID)
			}
		}
	}
}

func (c *client) AddToHistory(i models.Item) {
	c.ItemHistory[i.ID] = struct{}{}
}

func (c *client) IsInHistory(i models.Item) bool {
	if _, ok := c.ItemHistory[i.ID]; ok {
		return true
	}
	return false
}
