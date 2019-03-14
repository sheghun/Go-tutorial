package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of Deck
// Which is a slice of string
type deck []string

// newDeck for creating a new deck
func newDeck() deck {
	// Initialize the deck
	cards := deck{}

	// Create the card suits and the card values
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" Of "+value)
		}
	}
	return cards
}

// For printing our cards
func (d deck) print() {
	// Iterate over the slice
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deal function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// For converting a deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

/**
 * For saving a deck
 * writes a string representation of the deck to a file
 */
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

/**
 * For retrieve a deck from file
 * Returns a new deck from th string representation of a deck
 * saved on a file
 */
func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	// Check if error exists
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Convert to string then to a slice of string
	string := strings.Split(string(byteSlice), ",")

	// Convert to a deck and return
	return deck(string)
}

/**
 * Shuffle function
 * For shuffling the deck
 * Returns the deck with a random position of the cards
 */
func (d deck) shuffle() {
	// Create a new source for our random numbers
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
