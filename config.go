package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	Irc struct {
		Nick   string
		Server string
	}

	Discord struct {
		Token string
	}

	Channels map[string]string
}

func parseConfig(filename string) (*config, error) {
	cfg := &config{}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if err := yaml.NewDecoder(f).Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
