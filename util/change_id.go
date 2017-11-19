package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	changeIDURL string = "http://poe.ninja/api/Data/GetStats"
)

// NinjaStats stores the NextChangeID from the poe.ninja API.
type NinjaStats struct {
	NextChangeID string `json:"next_change_id"`
}

// GetLatestChangeID retrieves the latest stash fro poe.ninja.
func GetLatestChangeID() string {
	resp, err := http.Get(changeIDURL)
	if err != nil {
		return ""
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	var stats NinjaStats
	err = json.Unmarshal(b, &stats)
	if err != nil {
		return err.Error()
	}

	return strings.TrimSpace(stats.NextChangeID)
}
