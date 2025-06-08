# Pokedex CLI

A fully functional command-line Pokedex built in Go.  
It uses the PokeAPI to let users explore location areas, catch Pokémon, and inspect their stats.

## Features

- `help` — Show available commands
- `map` / `mapb` — Navigate through Pokémon locations
- `explore <location>` — Show all Pokémon in a location
- `catch <pokemon>` — Try to catch a Pokémon
- `inspect <pokemon>` — View stats of a caught Pokémon
- In-memory caching for faster repeated API requests

## Run It

```bash
go run .
