package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Spountil/bootdev_pokedexcli/internal/pokeapi"
	"github.com/Spountil/bootdev_pokedexcli/internal/pokecache"
)

func main() {
	const cacheTime = 1000 * time.Millisecond
	conf := pokeapi.Config{
		Commands: getCommands(),
		Cache:    pokecache.NewCache(cacheTime),
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")

	for {
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)

		cmd, ok := conf.Commands[words[0]]

		if len(words) > 1 {
			conf.Location = &words[1]
		}

		if ok {
			err := cmd.Callback(&conf)
			if err != nil {
				fmt.Println("Error: ", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
