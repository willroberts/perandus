package client

import (
	"errors"
	"time"
)

var (
	// The PoE API will issue temporary bans if requests are made more than once
	// per second.
	rateLimit = 1001 * time.Millisecond
)

// Poll sends at most one request per second to the PoE API, following the
// latest change ID to stay as close to real-time as possible.
func (c *client) Poll() error {
	rateLimiter := time.Tick(rateLimit)
	var exitError error

	for {
		<-rateLimiter

		stashes, err := c.getOne(c.NextChangeID)
		if err != nil {
			exitError = err
			break
		}
		if stashes.NextChangeID == "" {
			exitError = errors.New("empty change id")
			break
		}

		for _, s := range stashes.Stashes {
			c.FilterQueue <- s
		}

		c.NextChangeID = stashes.NextChangeID
	}

	return exitError
}
