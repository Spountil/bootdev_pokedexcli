package main

type cliCommand struct {
	name        string
	description string
	callback    func(conf *Config) error
}

type Config struct {
	commands map[string]cliCommand
	Next     *string
	Previous *string
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
