package filter

import (
	"io/ioutil"

	toml "github.com/pelletier/go-toml"
)

func readSettingsBytes() []byte {
	b, err := ioutil.ReadFile("settings.toml")
	if err != nil {
		return ""
	}
	return b
}

func parseSettings(b []byte) (*toml.Tree, error) {
	s, err := toml.Load(string(b))
	if err != nil {
		return &toml.Tree{}, err
	}
	_ = s
	// league := s.League, etc.
}
