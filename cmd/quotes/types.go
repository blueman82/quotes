package main

import (
	"errors"
	"math/rand"
)

// Quote represents a motivational quote with its author
type Quote struct {
	Text   string
	Author string
}

// ErrNoQuotes is returned when attempting to select from an empty quote list
var ErrNoQuotes = errors.New("no quotes available")

// SelectRandom returns a random quote from the provided slice using the given seed
// for reproducible randomness. Returns ErrNoQuotes if the quotes slice is empty.
func SelectRandom(quotes []Quote, seed int64) (Quote, error) {
	if len(quotes) == 0 {
		return Quote{}, ErrNoQuotes
	}

	source := rand.NewSource(seed)
	rng := rand.New(source)
	index := rng.Intn(len(quotes))

	return quotes[index], nil
}
