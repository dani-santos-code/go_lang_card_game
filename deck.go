package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)
// Create a new type of 
// Deck which is a slice of strings

type deck []string

// this borrows the behavior of a slice of strings
//like extends in OOP

// receiver
// by CONVENTION, the receiver is referred with
// a one or two-letter abbreviation that matches the
// type of the receiver
// because our receiver is of type deck, we use "d" (the
//first letter of deck)

func newDeck() deck {
	cards := deck {}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Jack", "Queen", "King"}

	for _, suit := range(cardSuits) {
		for _, value := range(cardValues){
			cards = append(cards, value+ " of "+suit)
		}
	}
	return cards
}

func (d deck) print(){
	for i, card := range d {
		fmt.Println(i, card)
	} 
}

//returns multiple values from one function
// returns two values of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
} 

func (d deck) toString() string {
	// converting a deck type into a slice of strings (converting back to what it was)
	return strings.Join([] string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	// error handling
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}