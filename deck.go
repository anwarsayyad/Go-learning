package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
// which is slice of strings
type deck [] string

func newDeck() deck {
	cards := deck{}

	cardsSuits := [] string {
		"Spades",
		"Diamonds",
		"Hearts",
		"Clubs",
	}
	cardValues := []string{"Ace", "two", "Three", "Four"}

	// used underscore to make error(declared but not used) go away
	for _, suits := range cardsSuits{
		for _, value := range cardValues{
			cards = append(cards, suits+ " of "+value)
		}
	}

	return cards
}

// this is reciver in GO and can be called on the any variable
// that is type of deck that we created
func (d deck) print () {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// in go we can return multiple values from on function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string (d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// option #1 - log the error and return a call to newDeck()
		// option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)

}


func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Generate a new seed from the current time (in nanoseconds).
	// Creating a new source with this seed ensures we get a different source
	// each time, producing truly random numbers.
	
	for i:= range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
