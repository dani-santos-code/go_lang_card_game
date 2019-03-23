package main

import "fmt"
// Create a new type of 
// Deck which is a slice of strings

type deck []string

// this borrows the behavior of a slice of strings
//like extends in OOP

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	} 
}