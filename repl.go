package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Spountil/bootdev_pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	var result []string
	words := strings.Fields(text)

	for _, word := range words {
		result = append(result, strings.ToLower(word))
	}
	return result
}

func getCommands() map[string]pokeapi.CliCommand {
	return map[string]pokeapi.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of 20 location areas in the Pokemon world, display the next 20 when called again",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the names of the 20 previous location ares in the Pokemon world when map was already called",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Return a list of all the pokemon in a location",
			Callback:    commandExplore,
		},
	}
}

func commandExit(conf *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *pokeapi.Config) error {
	fmt.Println("Usage:")

	for _, cmd := range conf.Commands {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}

	return nil
}

func commandMap(conf *pokeapi.Config) error {

	var url string

	if conf.Next != nil {
		url = *conf.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	var locResp pokeapi.LocationResponse

	data, err := pokeapi.FetchLocationAreas(conf, url)
	if err != nil {
		return err
	}

	locResp, err = pokeapi.MapResponse(data)
	results := locResp.Results
	conf.Next = locResp.Next
	conf.Previous = locResp.Previous

	for _, loc := range results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(conf *pokeapi.Config) error {

	var url string

	if conf.Previous != nil {
		url = *conf.Previous
	} else {
		fmt.Println("No previous map to query, querying the first 20 locations")
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	var locResp pokeapi.LocationResponse

	data, err := pokeapi.FetchLocationAreas(conf, url)
	if err != nil {
		return err
	}

	locResp, err = pokeapi.MapResponse(data)
	results := locResp.Results
	conf.Next = locResp.Next
	conf.Previous = locResp.Previous

	for _, loc := range results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandExplore(conf *pokeapi.Config) error {

	var url string

	fmt.Printf("Exploring %s...\n", *conf.Location)

	if len(*conf.Location) > 0 {
		url = "https://pokeapi.co/api/v2/location-area/" + *conf.Location
	} else {
		fmt.Println("Missing location name, try again.")
		return nil
	}

	var locResp pokeapi.LocationAreaDetails

	data, err := pokeapi.FetchLocationAreas(conf, url)
	if err != nil {
		return err
	}

	locResp, err = pokeapi.ExploreResponse(data)
	results := locResp.PokemonEncounters

	fmt.Println("Found Pokemon:")
	for _, pokemon := range results {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
