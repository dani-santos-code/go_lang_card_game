package main

import "fmt"
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
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven"}

	for _, suit := range(cardSuits) {
		for _, value := range(cardValues){
			cards = append(cards, value+ " of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	} 
}