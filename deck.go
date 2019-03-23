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

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	} 
}