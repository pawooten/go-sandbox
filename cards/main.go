package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" // declared type
	// card := "Ace of Spades" // inferred type, but := is only used for declaration + initialization
	// card := newCard() // inferred type by return type of function
	// fmt.Println((card))

	cards := []string{newCard(), "Ace of Diamonds", newCard()} // create a slice, which allows appending (arrays are fixed length)
	cards = append(cards, "Six of Spades")                     // returns a new slice (does not modify original slice)
	// fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
