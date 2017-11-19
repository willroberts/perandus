package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (c *client) getOne(changeID string) (*StashesResponse, error) {
	url := fmt.Sprintf("%s?id=%s", BaseURL, changeID)

	start := time.Now()
	resp, err := http.Get(url)
	latency := time.Since(start)
	if err != nil {
		return &StashesResponse{}, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &StashesResponse{}, err
	}

	var s StashesResponse
	if err := json.Unmarshal(b, &s); err != nil {
		return &StashesResponse{}, err
	}

	log.Println("statsd: request latency:", latency)

	return &s, nil
}
