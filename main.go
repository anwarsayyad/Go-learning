package main

import "fmt"




func main() {
	cards := newDeckFromFile("my cards")
	cards.print()
	cards.shuffle()
	fmt.Println("____After the Shuffle_____")
	cards.print()

	// cards = append(cards, "six of spades")
	// for i:=0; i<len(cards); i++ {
	// 	fmt.Printf(cards[i])
	// }


	// type convresions in thge go
	//[] byte ("Hi there!")
	// type we want --> [] byte
	// value we have --> ("Hi there!")

	// greeting := "Hi There"
	// fmt.Println([]byte(greeting))
}
