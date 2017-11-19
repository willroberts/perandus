package main

import (
	"log"

	"github.com/willroberts/perandus/client"
	"github.com/willroberts/perandus/util"
)

func main() {
	// Get the latest change ID.
	latest := util.GetLatestChangeID()
	log.Printf("Starting from change ID %s", latest)

	c := client.New(latest)
	c.Poll()
}
