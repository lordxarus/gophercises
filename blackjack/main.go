package main

import (
	"fmt"
	deck "gophercises/card-deck"
	"strings"
)

type Hand []deck.Card

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
		if c.Rank == deck.Ace {
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

func main() {
	var card deck.Card
	deck := deck.New(deck.Clone(3), deck.Shuffle)
	var player, dealer Hand

	for i := 0; i < 2; i++ {
		// using pointers because if we didn't
		// the Hand slice would contain copies of player
		// and dealer
		for _, hand := range []*Hand{&player, &dealer} {
			card, deck = draw(deck)
			*hand = append(*hand, card)
		}
	}

	var input string
	for input != "s" {
		fmt.Println("Player: ", player)
		fmt.Println("Dealer: ", dealer.DealerString())
		fmt.Println("What will do you do? (h)it, (s)tand?")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, deck = draw(deck)
			player = append(player, card)
		}
	}
	// If dScore <= 16 we (h)it
	// If dealer has a soft 17, we hit
	// (a dealer has a score of 17 when using an ace as 11)
	for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		card, deck = draw(deck)
		dealer = append(dealer, card)
	}
	pScore, dScore := player.Score(), dealer.Score()

	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player: ", player, "\nScore:", pScore)
	fmt.Println("Dealer: ", dealer, "\nScore:", dScore)

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
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
