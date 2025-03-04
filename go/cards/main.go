package main

func main() {
	cards := newDeck()
	cards.print()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	cards.writeToFile("my_cards")

	newCards := newDeckFromFile("my_cards")
	newCards.print()
}
