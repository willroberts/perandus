package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	ChangeIDURL string = "http://poe.ninja/api/Data/GetStats"
)

type NinjaStats struct {
	NextChangeID string `json:"next_change_id"`
}

func GetLatestChangeID() string {
	resp, err := http.Get(ChangeIDURL)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

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
