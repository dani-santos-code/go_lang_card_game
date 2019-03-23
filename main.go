package main

func main() {
	cards := deck {"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	fruits := deck {"Apple", "Banana", "Orange"}
	cards.print()
	fruits.print()
}

func newCard() string {
	return "Five of Diamonds"
}