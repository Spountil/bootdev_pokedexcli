# Pokedex CLI in Go

This is a REPL-based Pokedex command-line application built in Go. This project is part of the "Build a Pokedex in Go" course on [Boot.dev](https://boot.dev).

It uses the [PokéAPI](https://pokeapi.co/) to fetch data about Pokemon and their locations.

## Features

- Interactive REPL (Read-Eval-Print Loop) interface
- Fetch and display Pokemon location areas from the PokéAPI
- Navigate through location areas using `map` and `mapb` commands
- Explore areas to find Pokemon using the `explore` command
- Catch Pokemon using the `catch` command (catch rate depends on the Pokemon's base experience)
- Inspect caught Pokemon using the `inspect` command
- View all caught Pokemon using the `pokedex` command
- In-memory caching to prevent redundant API requests and improve performance

## Available Commands

- `help`: Displays a help message with all available commands.
- `exit`: Exits the Pokedex application.
- `map`: Displays the names of 20 location areas in the Pokemon world. Each subsequent call displays the next 20 locations.
- `mapb`: Displays the previous 20 locations.
- `explore <location_name>`: Explores a specific location to see a list of all the Pokemon that can be found in that area.
- `catch <pokemon_name>`: Attempts to catch a Pokemon. If successful, the Pokemon is added to your Pokedex.
- `inspect <pokemon_name>`: Inspects a Pokemon that you have caught to see its height, weight, stats, and types.
- `pokedex`: Lists all the Pokemon you have successfully caught.

## Getting Started

### Prerequisites

- Go 1.20 or higher installed on your machine.

### Installation & Usage

1. Clone the repository and navigate to the project directory:
   ```bash
   cd bootdev_pokedexcli
   ```

2. Build the project:
   ```bash
   go build -o pokedex
   ```

3. Run the executable:
   ```bash
   ./pokedex
   ```

## Acknowledgements

- [Boot.dev](https://boot.dev) for the excellent guided project course.
- [PokéAPI](https://pokeapi.co/) for the comprehensive Pokemon data API.
