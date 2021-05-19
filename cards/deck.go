package main

import "fmt"

// Create a new type of deck, which is a slice of strings
type deck []string

// This is a receiver function. (d deck) is the receiver of the function.
// Any variable of type deck now gets access to the print method.
// d is the instance variable reference, similar to this / self in java script
// By convention, Go developers never use this / self though
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Since this function creates a new deck, it doesn't make sense
// to specify the deck type as it's receiver.
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
