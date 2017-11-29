package main

import (
	"log"

	"github.com/willroberts/perandus/client"
	"github.com/willroberts/perandus/filter"
	"github.com/willroberts/perandus/util"
)

func main() {
	latest := util.GetLatestChangeID()
	log.Printf("Starting from change ID %s", latest)

	c := client.New(latest)
	itemCh, errCh := c.Poll()

	f, err := filter.New()
	if err != nil {
		log.Fatal("error creating filter:", err.Error())
	}

	filteredCh := f.FilterItems(itemCh)

	for {
		select {
		case i := <-filteredCh:
			log.Println("ALERT: Found matching item:")
			log.Println("\tName:", i.Name)
			log.Println("\tPrice:", i.Note)
			log.Println("\tSeller:", i.CharacterName)
			log.Println("\tID:", i.ID)
		case err := <-errCh:
			log.Fatal("error during polling:", err.Error())
		}
	}
}
