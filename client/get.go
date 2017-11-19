package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/willroberts/perandus/models"
)

const (
	baseURL string = "http://www.pathofexile.com/api/public-stash-tabs"
)

type StashesResponse struct {
	NextChangeID string `json:"next_change_id"`
	Stashes      []models.Stash
}

func (c *client) getOne(changeID string) (*StashesResponse, error) {
	url := fmt.Sprintf("%s?id=%s", baseURL, changeID)
	resp, err := http.Get(url)
	if err != nil {
		return &StashesResponse{}, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &StashesResponse{}, err
	}

	var s StashesResponse
	if err := json.Unmarshal(b, &s); err != nil {
		return &StashesResponse{}, err
	}

	return &s, nil
}
