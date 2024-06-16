package main

import "fmt"
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