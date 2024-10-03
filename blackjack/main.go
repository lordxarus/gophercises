// package main

// import (
// 	"fmt"
// 	deck "gophercises/card-deck"
// 	"strings"
// )

// type Hand []deck.Card

// func (h Hand) String() string {
// 	strs := make([]string, len(h))
// 	for i := range h {
// 		strs[i] = h[i].String()
// 	}
// 	return strings.Join(strs, ", ")
// }

// func (h Hand) DealerString() string {
// 	return h[0].String() + ", **HIDDEN**"
// }

// func Shuffle(gs GameState) GameState {
// 	ret := clone(gs)
// 	ret.Deck = deck.New(deck.Clone(3), deck.Shuffle)
// 	return ret
// }

// func Hit(gs GameState) GameState {
// 	ret := clone(gs)
// 	hand := ret.CurrentPlayer()
// 	var card deck.Card
// 	card, ret.Deck = draw(ret.Deck)
// 	*hand = append(*hand, card)
// 	if hand.Score() > 21 {
// 		return Stand(ret)
// 	}
// 	return ret
// }

// func Stand(gs GameState) GameState {
// 	ret := clone(gs)
// 	ret.State++
// 	return ret
// }

// func main() {
// 	// var card deck.Card
// 	var gs GameState
// 	gs = Shuffle(gs)
// 	for i := 0; i < 10; i++ {
// 		gs = Deal(gs)

// 		// If dScore <= 16 we (h)it
// 		// If dealer has a soft 17, we hit
// 		// (if we have an ace and are using it as an 11
// 		// and we have 17 as our score we have a soft 17

// 		gs = EndHand(gs)

// 	}
// }

// func clone(gs GameState) GameState {
// 	ret := GameState{
// 		State:  gs.State,
// 		Deck:   make([]deck.Card, len(gs.Deck)),
// 		Player: make(Hand, len(gs.Player)),
// 		Dealer: make(Hand, len(gs.Dealer)),
// 	}
// 	copy(ret.Deck, gs.Deck)
// 	copy(ret.Player, gs.Player)
// 	copy(ret.Dealer, gs.Dealer)

// 	return ret
// }

package main

import (
	"fmt"
	"gophercises/blackjack/blackjack"
)

func main() {
	game := blackjack.New()
	winnings := game.Play(blackjack.HumanPlayer())
	fmt.Println(winnings)
}
