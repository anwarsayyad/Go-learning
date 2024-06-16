package main

// import "fmt"


func main() {
	cards := newDeck()
	cards = append(cards, "six of spades")
	// for i:=0; i<len(cards); i++ {
	// 	fmt.Printf(cards[i])
	// }

	cards.print()
}
