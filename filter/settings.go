package filter

import (
	"errors"
	"io/ioutil"

	toml "github.com/pelletier/go-toml"
)

func (f *filter) parseSettings() error {
	b, err := ioutil.ReadFile("settings.toml")
	if err != nil {
		return err
	}

	s, err := toml.Load(string(b))
	if err != nil {
		return err
	}

	if err := validateSettings(s); err != nil {
		return err
	}

	f.League = s.Get("settings.league").(string)
	f.ItemName = s.Get("settings.item_name").(string)

	return nil
}

func validateSettings(settings *toml.Tree) error {
	if _, ok := settings.Get("settings.league").(string); !ok {
		return errors.New("invalid league")
	}

	if _, ok := settings.Get("settings.item_name").(string); !ok {
		return errors.New("invalid item name")
	}

	return nil
}
