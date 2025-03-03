package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	deckString := strings.Join([]string(d), ",")
	return deckString
}

func (d deck) writeToFile(filename string) {
	// ioutil
	bytes := []byte(d.toString())

	err := os.WriteFile(filename, bytes, 0666)

	if err != nil {
		log.Fatal(err)
	}

}

func newDeckFromFile(filename string) deck {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	deckString := string(bytes)
	cards := strings.Split(deckString, ",")
	return deck(cards)
}
