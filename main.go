package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Sundog28/pokedex-cli/internal/pokeapi"
	"github.com/Sundog28/pokedex-cli/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &pokeapi.Config{
		NextURL:        "",
		PrevURL:        "",
		Cache:          pokecache.NewMemoryCache(),
		CaughtPokemons: make(map[string]pokeapi.Pokemon),
	}

	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		args := strings.Fields(strings.ToLower(input))
		cmd := args[0]

		switch cmd {
		case "help":
			printHelp()
		case "exit":
			fmt.Println("Closing the Pokedex... Goodbye!")
			return
		case "map":
			if err := pokeapi.MapCommand(cfg); err != nil {
				fmt.Println("Error:", err)
			}
		case "mapb":
			if err := pokeapi.MapBackCommand(cfg); err != nil {
				fmt.Println("Error:", err)
			}
		case "explore":
			if len(args) < 2 {
				fmt.Println("Usage: explore <location-area-name>")
				continue
			}
			if err := pokeapi.ExploreCommand(cfg, args[1]); err != nil {
				fmt.Println("Error:", err)
			}
		case "catch":
			if len(args) < 2 {
				fmt.Println("Usage: catch <pokemon-name>")
				continue
			}
			if err := pokeapi.CatchCommand(cfg, args[1]); err != nil {
				fmt.Println("Error:", err)
			}
		case "inspect":
			if len(args) < 2 {
				fmt.Println("Usage: inspect <pokemon-name>")
				continue
			}
			pokeapi.InspectCommand(cfg, args[1])
		default:
			fmt.Println("Unknown command. Type 'help' for available commands.")
		}
	}
}

func printHelp() {
	fmt.Println(`
Available commands:
  help      - Display available commands
  exit      - Exit the Pokedex
  map       - View next 20 location areas
  mapb      - View previous 20 location areas
  explore   - Explore a location area (usage: explore <area_name>)
  catch     - Attempt to catch a Pokemon (usage: catch <pokemon_name>)
  inspect   - Show details of a caught Pokemon (usage: inspect <pokemon_name>)
`)
}
