package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardTypes := []string{"Heart", "Clover", "Spade", "Ace"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four"}

	for _, cardType := range cardTypes {
		for _, cardValue := range cardValues {
			cards = append(cards, cardType+" Of "+cardValue)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
 
func (d deck) toString() string {
	return strings.Join([]string(d),",")
}
