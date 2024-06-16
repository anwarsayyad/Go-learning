package main

// import "fmt"


func main() {
	cards := deck {"Ace of dimonds", newCard()}
	cards = append(cards, "six of spades")
	// for i:=0; i<len(cards); i++ {
	// 	fmt.Printf(cards[i])
	// }

	cards.print()
}

func newCard() string {
	return "Five diamonds"
}
