package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
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

	fmt.Println("----")
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

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
