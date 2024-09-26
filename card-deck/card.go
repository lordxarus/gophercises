//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"sort"
	"time"

	"math/rand"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Heart
	Club
	Joker // this is a special case
)

var suits = [...]Suit{Spade, Diamond, Heart, Club}

type Rank uint8

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

// functional options, or as I prefer to
// call them. Functional operators

func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func DefaultSort(cards []Card) []Card {
	// being able to modify parameters is
	// whacky!!
	sort.Slice(cards, Less(cards))
	return cards
}

// Create a sorter based with lessFn logic
func Sort(lessFn func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		// We capture lessFn with the closure
		// This part looks identical to DefaultSort
		// we just parameterized the lessFn
		sort.Slice(cards, lessFn(cards))
		return cards
	}
}

// i and j are indices into cards
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

// this is basically like converting from 2d coords (rank, suit) to 1d coords (absRank)
// (x, y) -> i
func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

// John makes this a package level variable so we can just set
// it from our tests
var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

// We're not shuffling inline, like Sort(), for simplicity's sake
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	// Perm(5)
	// [0, 4, 2, 1, 3]

	// We use these elements to index into ret slice.
	// And set ret[i] to the card value at the
	// idx j of the cards slice
	perm := shuffleRand.Perm(len(cards))
	for i, j := range perm {
		ret[i] = cards[j]
	}

	return ret

}

//
/*

If we just had:

func Jokers(cards []Card) []Card

we would have to use it like New(Joker, Joker, Joker)
to create three jokers. Instead we just return a function that
loops over n


*/
func Jokers(n int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Rank: Rank(i),
				Suit: Joker,
			})
		}
		return cards
	}
}

func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for _, c := range cards {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}

func Deck(n int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}
