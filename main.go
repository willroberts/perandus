package main

import (
	"log"

	"github.com/willroberts/perandus/client"
	"github.com/willroberts/perandus/util"
)

func main() {
	latest := util.GetLatestChangeID()
	log.Printf("Starting from change ID %s", latest)
	c := client.New(latest)
	c.Poll()
}
