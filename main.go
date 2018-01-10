package main

func main() {
	//Save Deck
	//cards := newDeck()
	//cards.saveToFile("my_cards")

	//Read Deck
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shufffle()
	cards.print()
}
