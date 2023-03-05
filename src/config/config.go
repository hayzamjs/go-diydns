package config

import (
	"encoding/json"
	"io/ioutil"
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
)

type Config struct {
	Name string `json:"name"`
	Provider string `json:"provider"`
	Token string `json:"token"`
	Domain string `json:"domain"`
	Interval int `json:"interval"`
}

var supportedProviders = [...]string{"cloudflare"}


func ReadConfig(path string) ([]Config, error) {
	var config []Config

	file, err := ioutil.ReadFile(path)

	if err != nil {
		return config, errors.New("Configuration at config.json not found")
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		return config, errors.New("Configuration at config.json is invalid")
	}

	for _, c := range config {
		var errStr string = ""

		if c.Name == "" {
			errStr = fmt.Sprintf("Configuration at config.json is invalid, name is required for every entry")
			return config, errors.New(errStr)
		}

		if c.Provider == "" {
			errStr = fmt.Sprintf("Configuration at config.json is invalid, provider is required for %s", c.Name)
			return config, errors.New(errStr)
		} else {
			var found bool = false

			for _, p := range supportedProviders {
				if c.Provider == p {
					found = true
				}
			}

			if !found {
				errStr = fmt.Sprintf("Configuration at config.json is invalid, provider is not supported for %s", c.Name)
				return config, errors.New(errStr)
			}
		}

		if c.Token == "" {
			errStr = fmt.Sprintf("Configuration at config.json is invalid, token is required for %s", c.Name)
			return config, errors.New(errStr)
		}

		if c.Domain == "" {
			errStr = fmt.Sprintf("Configuration at config.json is invalid, domain is required for %s", c.Name)
			return config, errors.New(errStr)
		} else {
			if !govalidator.IsDNSName(c.Domain) {
				errStr = fmt.Sprintf("Configuration at config.json is invalid, domain is not valid for %s", c.Name)
				return config, errors.New(errStr)
			}
		}

		if c.Interval == 0 {
			errStr = fmt.Sprintf("Configuration at config.json is invalid, interval is required for %s", c.Name)
			return config, errors.New(errStr)
		}
	}

	return config, nil
}