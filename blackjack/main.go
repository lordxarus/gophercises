package main

import (
	"fmt"
	cd "gophercises/card-deck"
	"strings"
)

type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type Hand []cd.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == cd.Ace {
			return minScore + 10
		}
	}
	return minScore
}

func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
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

func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StateDealerTurn:
		return &gs.Dealer
	case StatePlayerTurn:
		return &gs.Player
	default:
		panic("it isn't currently any player's turn")
	}
}

func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	ret.Deck = cd.New(cd.Clone(3), cd.Shuffle)
	return ret
}

func Deal(gs GameState) GameState {
	ret := clone(gs)
	ret.Player = make(Hand, 0, 5)
	ret.Dealer = make(Hand, 0, 5)
	var card cd.Card
	for i := 0; i < 2; i++ {
		card, ret.Deck = draw(ret.Deck)
		ret.Player = append(ret.Player, card)
		card, ret.Deck = draw(ret.Deck)
		ret.Dealer = append(ret.Dealer, card)
	}
	ret.State = StatePlayerTurn
	return ret
}

func Hit(gs GameState) GameState {
	ret := clone(gs)
	hand := ret.CurrentPlayer()
	var card cd.Card
	card, ret.Deck = draw(ret.Deck)
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(ret)
	}
	return ret
}

func Stand(gs GameState) GameState {
	ret := clone(gs)
	ret.State++
	return ret
}

func EndHand(gs GameState) GameState {
	ret := clone(gs)
	pScore, dScore := ret.Player.Score(), ret.Dealer.Score()

	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player: ", ret.Player, "\nScore:", pScore)
	fmt.Println("Dealer: ", ret.Dealer, "\nScore:", dScore)
	fmt.Println()

	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("Dealer busted")
	case pScore > dScore:
		fmt.Println("You win!")
	case dScore > pScore:
		fmt.Println("You lose")
	case dScore == pScore:
		fmt.Println("Draw")
	}
	ret.Player = nil
	ret.Dealer = nil
	return ret
}

func main() {
	// var card deck.Card
	var gs GameState
	gs = Shuffle(gs)
	for i := 0; i < 10; i++ {
		gs = Deal(gs)

		var input string
		for gs.State == StatePlayerTurn {
			fmt.Println("Player: ", gs.Player)
			fmt.Println("Dealer: ", gs.Dealer.DealerString())
			fmt.Println("What will do you do? (h)it, (s)tand?")
			fmt.Scanf("%s\n", &input)
			switch input {
			case "h":
				gs = Hit(gs)
			case "s":
				gs = Stand(gs)
			default:
				fmt.Println("invalid option: ", input)
			}
		}

		// If dScore <= 16 we (h)it
		// If dealer has a soft 17, we hit
		// (a dealer has a score of 17 when using an ace as 11)

		for gs.State == StateDealerTurn {
			if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
				gs = Hit(gs)
			} else {
				gs = Stand(gs)
			}
		}

		gs = EndHand(gs)

	}
}

func draw(cards []cd.Card) (cd.Card, []cd.Card) {
	return cards[0], cards[1:]
}

type GameState struct {
	Deck   []cd.Card
	State  State
	Player Hand
	Dealer Hand
}

func clone(gs GameState) GameState {
	ret := GameState{
		State:  gs.State,
		Deck:   make([]cd.Card, len(gs.Deck)),
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)

	return ret
}
