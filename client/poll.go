package client

import "time"

var (
	rateLimit = 1001 * time.Millisecond
)

func (c *client) Poll() {
	rateLimiter := time.Tick(rateLimit)
	for {
		<-rateLimiter
		//get(c.NextChangeID)
		//c.NextChangeID = ?
	}
}
