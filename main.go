package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card := newCard()
	// here, we are relying upon the Go compiler to figure out
	// that card is supposed to contain a string
	// the colon equal indicates that the compiler needs to figure out
	// the colon equals is only iused when you assign to a variable for 
	// the first time! it's not used for reassignemnt!
	// to reassign we would go for:
	// card = "Five of Spades"
	// fmt.Println(card)
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	} 


	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}