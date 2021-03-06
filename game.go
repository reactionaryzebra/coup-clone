package main

import (
	"fmt"
	"math/rand"
	"time"
)

// type game struct {
// 	deck    Deck
// 	players []player
// }

type deck struct {
	initialCards []string
	currentCards []string
}

func (d *deck) draw(num int) []string {
	rand.Seed(time.Now().UnixNano())
	cards := make([]string, 0)
	for i := 0; i < num; i++ {
		index := rand.Intn(len(d.currentCards))
		card := d.currentCards[index]
		cards = append(cards, card)
		d.currentCards = append(d.currentCards[:index], d.currentCards[index+1:]...)
	}
	return cards
}

var cards = []string{"duke", "duke", "duke", "captain", "captain", "captain", "ambassador", "ambassador", "ambassador", "assassin", "assassin", "assassin", "contessa", "contessa", "contessa"}

var actionToCharMap = map[string]string{
	"tax":         "duke",
	"assassinate": "assassin",
	"steal":       "captain",
	"exchange":    "ambassador",
	"income":      "any",
	"foreignAid":  "any",
	"coup":        "any",
}
var blockActionToCharMap = map[string]string{
	"foreignAid":      "duke",
	"assassination":   "contessa",
	"stealCaptain":    "captain",
	"stealAmbassador": "ambassador",
}

type player struct {
	cards []string
	coins int
}

func (p *player) changeCoins(num int) int {
	p.coins += num
	return p.coins
}

func (p *player) removeCard(index int) []string {
	p.cards = append(p.cards[:index], p.cards[index+1:]...)
	return p.cards
}

func (p *player) addCard(newCard string) []string {
	p.cards = append(p.cards, newCard)
	return p.cards
}

func (p *player) checkTruth(character string) bool {
	for _, char := range p.cards {
		if char == character {
			return true
		}
	}
	return false
}

type game struct {
	numPlayers int
	players    []player
	turnIndex  int
	deck       deck
}

func (g *game) init() {
	g.deck = deck{initialCards: cards, currentCards: cards}
	for i := 0; i < g.numPlayers; i++ {
		g.players = append(g.players, player{cards: g.deck.draw(2), coins: 2})
	}
}

func (g *game) passTurn() int {
	var newIndex int
	if len(g.players)-1 == g.turnIndex {
		newIndex = 0
	} else {
		newIndex = g.turnIndex + 1
	}
	g.turnIndex = newIndex
	return newIndex
}

func (g *game) removePlayer(position int) {
	g.players = append(g.players[:position], g.players[position+1:]...)
}

func main() {
	g := game{numPlayers: 4}
	g.init()
	fmt.Println(g.players)
	fmt.Println(g.players[0].checkTruth("contessa"))
	g.removePlayer(1)
	fmt.Println(g.players)
}
