package client

import (
	"log"

	"github.com/willroberts/perandus/filter"
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

			if c.Filter.Matches(i) && i.Note != "" {
				log.Println("ALERT! Found matching item:")
				log.Println("\tName:", i.Name)
				log.Println("\tPrice:", i.Note)
				log.Println("\tSeller:", s.LastCharacterName)
			}
		}
	}
}
