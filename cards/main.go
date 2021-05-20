package main

func main() {
	// var card string = "Ace of Spades" // declared type
	// card := "Ace of Spades" // inferred type, but := is only used for declaration + initialization
	// card := newCard() // inferred type by return type of function
	// fmt.Println((card))

	// 	cards := deck{newCard(), "Ace of Diamonds", newCard()} // create a slice, which allows appending (arrays are fixed length)
	// 	cards = append(cards, "Six of Spades")                 // returns a new slice (does not modify original slice)
	// 	// fmt.Println(cards)

	// 	cards.print()

	// cards := newDeck()
	// // cards.print()
	// hand, remainingDeck := deal(cards, 5) // handSize = 5

	// fmt.Println("Your Hand")
	// hand.print()
	// fmt.Println("\n\nRemaining Deck")
	// remainingDeck.print()

	// greeting := "Hi There!"
	// // Type conversion from string to byte slice
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// cards.saveToFile("cards.txt")
	// cards := newDeckFromFile("cards.txt")
	// cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
