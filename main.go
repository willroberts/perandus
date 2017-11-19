package main

import (
	"log"

	"github.com/willroberts/perandus/client"
	"github.com/willroberts/perandus/util"
)

func main() {
	latest := util.GetLatestChangeID()
	log.Printf("Starting from change ID %s", latest)
	c, err := client.New(latest)
	if err != nil {
		log.Fatal("Failed to create client:", err.Error())
	}
	if err := c.Poll(); err != nil {
		log.Fatal("Failed while polling:", err.Error())
	}
}
