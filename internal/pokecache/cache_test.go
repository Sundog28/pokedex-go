package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)
	cache.Add("https://pokeapi.co", []byte("pokemon"))

	val, ok := cache.Get("https://pokeapi.co")
	if !ok || string(val) != "pokemon" {
		t.Errorf("expected to retrieve cached value")
	}
}

func TestReapLoop(t *testing.T) {
	const base = 10 * time.Millisecond
	cache := NewCache(base)

	cache.Add("key", []byte("value"))

	time.Sleep(base + 10*time.Millisecond)

	_, ok := cache.Get("key")
	if ok {
		t.Errorf("expected cache entry to be reaped")
	}
}
