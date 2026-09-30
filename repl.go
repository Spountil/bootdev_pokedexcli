package main

import (
	"fmt"
	"math/rand"
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
		"catch": {
			Name:        "catch",
			Description: "Function to catch a Pokemon passed following the command",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Return the stats of caught Pokemons",
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Return the Pokemon in the Pokedex",
			Callback:    commandPokedex,
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

	data, err := pokeapi.FetchApi(conf, url)
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

	data, err := pokeapi.FetchApi(conf, url)
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

	fmt.Printf("Exploring %s...\n", *conf.Param)

	if len(*conf.Param) > 0 {
		url = "https://pokeapi.co/api/v2/location-area/" + *conf.Param
	} else {
		fmt.Println("Missing location name, try again.")
		return nil
	}

	var locResp pokeapi.LocationAreaDetails

	data, err := pokeapi.FetchApi(conf, url)
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

func commandCatch(conf *pokeapi.Config) error {

	_, ok := conf.Pokedex[*conf.Param]

	if ok {
		fmt.Printf("%s already in the Pokedex\n", *conf.Param)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", *conf.Param)

	url := "https://pokeapi.co/api/v2/pokemon/" + *conf.Param

	var locResp pokeapi.PokemonDetails

	data, err := pokeapi.FetchApi(conf, url)
	if err != nil {
		return err
	}

	locResp, err = pokeapi.PokemonResponse(data)
	// the ceiling of the base exp of a Pokemon is 635, so putting the ceiling at 700 makes sens
	baseExpCeiling := 700
	roll := rand.Intn(baseExpCeiling)
	baseExp := locResp.BaseExperience

	if roll > baseExp {
		fmt.Printf("%s was caught!\n", *conf.Param)
		conf.Pokedex[*conf.Param] = locResp
	} else {
		fmt.Printf("%s escaped!\n", *conf.Param)
	}

	return nil
}

func commandInspect(conf *pokeapi.Config) error {
	result, ok := conf.Pokedex[*conf.Param]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	stats := struct {
		hp             int
		attack         int
		defense        int
		specialAttack  int
		specialDefense int
		speed          int
	}{}

	for _, stat := range result.Stats {
		switch statType := stat.Stat; statType.Name {
		case "hp":
			stats.hp = stat.BaseStat
		case "attack":
			stats.attack = stat.BaseStat
		case "defense":
			stats.defense = stat.BaseStat
		case "special-attack":
			stats.specialAttack = stat.BaseStat
		case "special-defense":
			stats.specialDefense = stat.BaseStat
		case "speed":
			stats.speed = stat.BaseStat
		}
	}

	fmt.Printf("Name: %s\n", result.Name)
	fmt.Printf("Height: %d\n", result.Height)
	fmt.Printf("Weight: %d\n", result.Weight)
	fmt.Println("Stats:")
	fmt.Printf(" -hp: %d\n", stats.hp)
	fmt.Printf(" -attack: %d\n", stats.attack)
	fmt.Printf(" -defense: %d\n", stats.defense)
	fmt.Printf(" -special-attack: %d\n", stats.specialAttack)
	fmt.Printf(" -special-defense: %d\n", stats.specialDefense)
	fmt.Printf(" -speed: %d\n", stats.speed)
	fmt.Println("Types:")
	for _, pokemonType := range result.Types {
		fmt.Printf(" -%s\n", pokemonType.Type.Name)
	}

	return nil
}

func commandPokedex(conf *pokeapi.Config) error {
	if len(conf.Pokedex) == 0 {
		fmt.Println("Your pokedex is empty, go catch some Pokemon!")
		return nil
	}

	fmt.Println("Your Pokedex")
	for _, pokemon := range conf.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
