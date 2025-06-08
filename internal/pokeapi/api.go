package pokeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/Sundog28/pokedex-cli/internal/pokecache"
)

func fetchURL(url string, cache pokecache.Cache) ([]byte, error) {
	if cache != nil {
		if data, ok := cache.Get(url); ok {
			return data, nil
		}
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if cache != nil {
		cache.Set(url, data)
	}

	return data, nil
}

func MapCommand(cfg *Config) error {
	url := cfg.NextURL
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	data, err := fetchURL(url, cfg.Cache)
	if err != nil {
		return err
	}

	var resp LocationAreaResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return err
	}
	for _, r := range resp.Results {
		fmt.Println(r.Name)
	}

	cfg.NextURL = resp.Next
	cfg.PrevURL = resp.Previous
	return nil
}

func MapBackCommand(cfg *Config) error {
	if cfg.PrevURL == "" {
		fmt.Println("No previous page.")
		return nil
	}
	data, err := fetchURL(cfg.PrevURL, cfg.Cache)
	if err != nil {
		return err
	}

	var resp LocationAreaResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return err
	}
	for _, r := range resp.Results {
		fmt.Println(r.Name)
	}

	cfg.NextURL = resp.Next
	cfg.PrevURL = resp.Previous
	return nil
}

func ExploreCommand(cfg *Config, area string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", area)
	data, err := fetchURL(url, cfg.Cache)
	if err != nil {
		return err
	}

	var detail LocationAreaDetail
	if err := json.Unmarshal(data, &detail); err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\nFound Pokemon:\n", area)
	for _, e := range detail.PokemonEncounters {
		fmt.Printf(" - %s\n", e.Pokemon.Name)
	}
	return nil
}

func CatchCommand(cfg *Config, name string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	data, err := fetchURL(url, cfg.Cache)
	if err != nil {
		return err
	}

	var p Pokemon
	if err := json.Unmarshal(data, &p); err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", p.Name)

	rand.Seed(time.Now().UnixNano())
	chance := 0.5 - float64(p.BaseExperience)/200.0
	if chance < 0.05 {
		chance = 0.05
	}

	if rand.Float64() < chance {
		cfg.CaughtPokemons[p.Name] = p
		fmt.Printf("%s was caught!\n", p.Name)
	} else {
		fmt.Printf("%s escaped!\n", p.Name)
	}
	return nil
}

func InspectCommand(cfg *Config, name string) {
	p, ok := cfg.CaughtPokemons[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", p.Name, p.Height, p.Weight)
	for _, s := range p.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
}
