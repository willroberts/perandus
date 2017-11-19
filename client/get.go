package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
