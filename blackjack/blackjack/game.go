package blackjack

import (
	"fmt"
	deck "gophercises/card-deck"
)

type state int8

const (
	statePlayerTurn state = iota
	stateDealerTurn
	stateHandOver
)

func New() Game {
	return Game{
		state:   statePlayerTurn,
		pDealer: dealerPlayer{},
		balance: 0,
	}
}

type Game struct {
	// unexported fields
	deck    []deck.Card
	state   state
	player  []deck.Card
	dealer  []deck.Card
	pDealer Player
	balance int
}

func (g *Game) currentHand() *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player
	case stateDealerTurn:
		return &g.dealer
	default:
		panic("it isn't currently any player's turn")
	}
}

func (g *Game) Play(p Player) int {
	g.deck = deck.New(deck.Clone(3), deck.Shuffle)

	for i := 0; i < 10; i++ {
		deal(g)
		for g.state == statePlayerTurn {
			// protect the hand
			hand := make([]deck.Card, len(g.player))
			copy(hand, g.player)
			move := p.Play(hand, g.dealer[0])
			move(g)
		}
		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.dealer))
			copy(hand, g.dealer)
			move := g.pDealer.Play(hand, g.dealer[0])
			move(g)
		}
		endHand(g, p)
	}
	return g.balance
}

type Move func(*Game)

func MoveHit(g *Game) {
	hand := g.currentHand()
	var card deck.Card
	card, g.deck = draw(g.deck)
	*hand = append(*hand, card)
	if Score(*hand...) > 21 {
		MoveStand(g)
	}
}

func MoveStand(g *Game) {
	g.state++
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func deal(g *Game) {
	g.player = make([]deck.Card, 0, 5)
	g.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.deck = draw(g.deck)
		g.player = append(g.player, card)
		card, g.deck = draw(g.deck)
		g.dealer = append(g.dealer, card)
	}
	g.state = statePlayerTurn
}

func endHand(g *Game, p Player) {
	pScore, dScore := Score(g.player...), Score(g.dealer...)
	// TODO: Figure out winnings and add / subtract them
	switch {
	case pScore > 21:
		fmt.Println("You busted")
		g.balance--
	case dScore > 21:
		fmt.Println("Dealer busted")
		g.balance++
	case pScore > dScore:
		fmt.Println("You win!")
		g.balance++
	case dScore > pScore:
		fmt.Println("You lose")
		g.balance--
	case dScore == pScore:
		fmt.Println("Draw")
	}
	fmt.Println()
	p.Results([][]deck.Card{g.player}, g.dealer)
	g.player = nil
	g.dealer = nil
}

// Score will take in a hand of cards and return the best blackjack
// score possible with the hand.
func Score(hand ...deck.Card) int {
	minScore := minScore(hand...)
	if minScore > 11 {
		return minScore
	}
	for _, c := range hand {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

// returns true if the score is 17 and we are using an ace
// as an 11
func IsSoft(hand ...deck.Card) bool {
	return minScore(hand...) != Score(hand...)
}

func minScore(hand ...deck.Card) int {
	score := 0
	for _, c := range hand {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
