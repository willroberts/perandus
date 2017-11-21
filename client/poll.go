package client

import (
	"errors"
	"time"

	"github.com/willroberts/perandus/models"
)

var (
	// The PoE API will issue temporary bans if requests are made more than once
	// per second.
	rateLimit = 1 * time.Second
)

// Poll sends at most one request per second to the PoE API, following the
// latest change ID to stay as close to real-time as possible.
func (c *client) Poll() (chan models.Item, chan error) {
	itemCh := make(chan models.Item)
	errCh := make(chan error, 1)

	go func() {
		rateLimiter := time.Tick(rateLimit)
		for {
			<-rateLimiter

			stashes, err := c.getOne(c.NextChangeID)
			if err != nil {
				errCh <- err
				break
			}

			if stashes.NextChangeID == "" {
				errCh <- errors.New("empty change ID")
				break
			}

			for _, s := range stashes.Stashes {
				for _, i := range s.Items {
					i.Name = models.StripLocalizationTags(i.Name)
					if i.Name == "" {
						i.Name = models.StripLocalizationTags(i.TypeLine)
					}
					i.CharacterName = s.LastCharacterName

					if !c.isInHistory(i) {
						c.addToHistory(i)
					} else {
						continue
					}

					itemCh <- i
				}
			}

			c.NextChangeID = stashes.NextChangeID
		}
	}()

	return itemCh, errCh
}
