package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Spountil/bootdev_pokedexcli/internal/pokecache"
)

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world, display the next 20 when called again",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the 20 previous location ares in the Pokemon world when map was already called",
			callback:    commandMapb,
		},
	}
}

func commandExit(conf *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *Config) error {
	fmt.Println("Usage:")

	for _, cmd := range conf.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func fetchLocationAreas(conf *Config, url string) error {
	var data []byte
	result, ok := conf.Cache.CacheMap[url]

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		conf.Cache.CacheMap[url] = pokecache.CacheEntry{
			CreatedAt: time.Now(),
			Val:       data,
		}

		if res.StatusCode > 299 {
			return fmt.Errorf("response's status not 200. Status: %d", res.StatusCode)
		}
	} else {
		data = result.Val
	}

	var locResp LocationResponse
	err := json.Unmarshal(data, &locResp)
	if err != nil {
		return err
	}

	results := locResp.Results
	conf.Next = locResp.Next
	conf.Previous = locResp.Previous

	for _, loc := range results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMap(conf *Config) error {

	var url string

	if conf.Next != nil {
		url = *conf.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	fetchLocationAreas(conf, url)

	return nil
}

func commandMapb(conf *Config) error {

	var url string

	if conf.Previous != nil {
		url = *conf.Previous
	} else {
		fmt.Println("No previous map to query, querying the first 20 locations")
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	fetchLocationAreas(conf, url)

	return nil
}

func main() {
	const cacheTime = 1000 * time.Millisecond
	conf := Config{
		commands: getCommands(),
		Cache:    pokecache.NewCache(cacheTime),
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")

	for {
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)

		cmd, ok := conf.commands[words[0]]

		if ok {
			err := cmd.callback(&conf)
			if err != nil {
				fmt.Println("Error: ", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
