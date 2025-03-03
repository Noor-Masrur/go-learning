package main

func main() {
	cards := newDeck()

	// cards.print()

	// first, second := deal(cards, 5)

	// first.print()
	// println("----")
	// second.print()

	cards.writeToFile("deck.txt")

	newDeck := newDeckFromFile("deck.txt")

	newDeck.print()
}
