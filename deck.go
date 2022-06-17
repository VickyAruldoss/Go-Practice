package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
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
	return strings.Join([]string(d), ",")
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
