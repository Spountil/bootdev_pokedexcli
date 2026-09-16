package main

import (
	"github.com/Spountil/bootdev_pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *Config) error
}

type Config struct {
	commands map[string]cliCommand
	Next     *string
	Previous *string
	Cache    *pokecache.Cache
}

type ShallowLocation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationResponse struct {
	Count    int               `json:"count"`
	Next     *string           `json:"next"`
	Previous *string           `json:"previous"`
	Results  []ShallowLocation `json:"results"`
}
