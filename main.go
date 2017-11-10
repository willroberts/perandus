package main

import (
	"io/ioutil"
	"log"
	"time"

	toml "github.com/pelletier/go-toml"
	"github.com/willroberts/perandus/client"
)

func main() {
	settings, err := toml.Load(readSettings())
	if err != nil {
		log.Fatal("failed to parse settings: ", err)
	}
	next := settings.Get("settings.next_change_id").(string)
	log.Println("Initial change ID: ", next)

	c := client.NewClient()

	start := time.Now()
	resp, err := c.GetStashes(next)
	if err != nil {
		log.Fatal("failed to retrieve stashes: ", err)
	}
	latency := time.Since(start)

	log.Println("Stashes:", len(resp.Stashes))
	log.Println("Next Change ID:", resp.NextChangeID)
	log.Println("Latency:", latency)
}

func readSettings() string {
	bytes, err := ioutil.ReadFile("settings.toml")
	if err != nil {
		return ""
	}
	return string(bytes)
}
