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

// defaultQuotes contains the built-in collection of quotes
var defaultQuotes = []Quote{
	{Text: "Be the change you wish to see in the world", Author: "Gandhi"},
	{Text: "Code is poetry", Author: "Unknown"},
	{Text: "The only way to do great work is to love what you do", Author: "Steve Jobs"},
	{Text: "Innovation distinguishes between a leader and a follower", Author: "Steve Jobs"},
	{Text: "Stay hungry, stay foolish", Author: "Steve Jobs"},
	{Text: "Simplicity is the ultimate sophistication", Author: "Leonardo da Vinci"},
	{Text: "Talk is cheap. Show me the code", Author: "Linus Torvalds"},
	{Text: "First, solve the problem. Then, write the code", Author: "John Johnson"},
	{Text: "Any fool can write code that a computer can understand. Good programmers write code that humans can understand", Author: "Martin Fowler"},
	{Text: "The best error message is the one that never shows up", Author: "Thomas Fuchs"},
}

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
