package client

import (
	"log"
	"regexp"
)

var setRemover = regexp.MustCompile("<<.*>>")

func (c *client) runFilterWorker() {
	for {
		s := <-c.FilterQueue

		if len(s.Items) == 0 {
			continue
		}

		log.Println("Items parsed:", len(s.Items))
		/*for _, i := range s.Items {
			if i.Name != "" {
				log.Println("Item:", pruneItemName(i.Name))
			}
		}*/
	}
}

func pruneItemName(name string) string {
	b := setRemover.ReplaceAll([]byte(name), []byte(""))
	return string(b)
}
