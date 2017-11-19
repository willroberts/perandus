package client

import (
	"log"
	"time"
)

var (
	rateLimit = 1001 * time.Millisecond
)

// Poll sends at most one request per second to the PoE API, following the
// latest change ID to stay as close to real-time as possible.
func (c *client) Poll() {
	rateLimiter := time.Tick(rateLimit)
	for {
		<-rateLimiter

		stashes, err := c.getOne(c.NextChangeID)
		if err != nil {
			log.Println("error:", err.Error())
		}
		if stashes.NextChangeID == "" {
			log.Println("empty change ID encountered")
			break
		}

		for _, s := range stashes.Stashes {
			c.FilterQueue <- s
		}

		c.NextChangeID = stashes.NextChangeID
	}
}
