package blackjack

import (
	"fmt"
	deck "gophercises/card-deck"
)

type Player interface {
	Bet() int
	Play(hand []deck.Card, dealer deck.Card) Move
	Results(hand [][]deck.Card, dealer []deck.Card)
}

type humanPlayer struct{}

func HumanPlayer() Player {
	return humanPlayer{}
}

type dealerPlayer struct{}

func (d dealerPlayer) Play(hand []deck.Card, dealer deck.Card) Move {
	dScore := Score(hand...)
	if dScore <= 16 || (dScore == 17 && IsSoft(hand...)) {
		return MoveHit
	} else {
		return MoveStand
	}
}

func (p dealerPlayer) Bet() int {
	return 1
}

func (p dealerPlayer) Results(hand [][]deck.Card, dealer []deck.Card) {
	// noop
}

func (p humanPlayer) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println("Player: ", hand)
		fmt.Println("Dealer: ", dealer)
		fmt.Println("What will do you do? (h)it, (s)tand?")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		default:
			fmt.Println("invalid option: ", input)
		}
	}

}

func (p humanPlayer) Bet() int {
	return 1
}

func (p humanPlayer) Results(hand [][]deck.Card, dealer []deck.Card) {
	fmt.Println("==FINAL HANDS==")
	// pScore, dScore := ret.Player.Score(), ret.Dealer.Score()

	fmt.Println("Player: ", hand)
	fmt.Println("Dealer: ", dealer)

}
