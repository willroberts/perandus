package client

import (
	"log"
)

func (c *client) runFilterWorker() {
	for {
		s := <-c.FilterQueue

		if len(s.Items) == 0 {
			continue
		}

		for _, i := range s.Items {
			if c.Filter.Matches(i) {
				log.Println("ALERT! Found matching item:", i)
			}
		}
	}
}
