package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" // declared type
	// card := "Ace of Spades" // inferred type, but := is only used for declaration + initialization
	card := newCard() // inferred type by return type of function
	fmt.Println((card))
}

func newCard() string {
	return "Five of Diamonds"
}
