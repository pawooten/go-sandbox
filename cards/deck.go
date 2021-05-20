package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// Slices are zero indexed.
// Slices can be index fruits[2]
// Go has subset functionality built into slices
// fruits[0:2] is fruits[startIndexIncluding : upToNotIncluding]
// fruits[0:2] and fruits[:2] are equivalent
// fruits[2:] means index at 2 to end of slice

// Go has support for returning multiple values from one function
// This function returns to values, both of type deck
// Return the dealt hand and the remaining deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// Return a string representation of a deck
func (d deck) toString() string {
	// Type conversion from deck to []string, then join []string with a comma separator
	return strings.Join([]string(d), ",")
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		// exit program entirely
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	// Convert the current time to int64 and use as the seed for the source of the random number generator
	randy := rand.New(rand.NewSource(time.Now().UnixNano()))

	for index := range d {
		newPosition := randy.Intn(len(d) - 1)

		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
