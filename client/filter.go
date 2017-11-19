package client

import (
	"github.com/willroberts/perandus/alert"
	"github.com/willroberts/perandus/models"
)

func (c *client) runFilterWorker() {
	for {
		s := <-c.FilterQueue

		if len(s.Items) == 0 {
			continue
		}

		for _, i := range s.Items {
			i.Name = models.StripLocalizationTags(i.Name)

			if !c.isInHistory(i) {
				c.addToHistory(i)
			} else {
				// We've already seen this item; don't alert.
				continue
			}

			if c.Filter.Matches(i) && i.Note != "" {
				alert.ConsoleLogAlert(i, s)
			}
		}
	}
}

func (c *client) addToHistory(i models.Item) {
	c.ItemHistory[i.ID] = struct{}{}
}

func (c *client) isInHistory(i models.Item) bool {
	if _, ok := c.ItemHistory[i.ID]; ok {
		return true
	}
	return false
}
