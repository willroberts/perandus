package main

import (
	"log"

	"github.com/willroberts/perandus/alert"
	"github.com/willroberts/perandus/client"
	"github.com/willroberts/perandus/util"
)

func main() {
	latest := util.GetLatestChangeID()
	log.Printf("Starting from change ID %s", latest)
	c := client.New(latest)
	itemCh, errCh := c.Poll()
	//create filter
	for {
		select {
		case i := <-itemCh:
			//if filter.matches() && i.Note != "" {
			if i.Note != "" {
				alert.ConsoleLogAlert(i)
			}
		case err := <-errCh:
			log.Fatal("error during polling:", err.Error())
		}
	}
}
