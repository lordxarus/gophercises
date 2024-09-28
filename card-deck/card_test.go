package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
	// Joker

}

func TestNew(t *testing.T) {
	cards := New()
	// 13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Errorf("wrong num of cards in a new deck: %d", len(cards))
	}
}

// I don't think these sorting tests actually test anything
// New() returns a sorted deck by default

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	// without these parenthesis we fail to compile
	if cards[0] != (Card{Rank: Ace, Suit: Spade}) {
		t.Error("wanted Ace of Spades as first card, but got", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	if cards[0] != (Card{Rank: Ace, Suit: Spade}) {
		t.Error("wanted Ace of Spades as first card, but got", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("expected 3 jokers, got", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(c Card) bool {
		return c.Rank == Two || c.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if filter(c) {
			t.Errorf("found: %s, but it was meant to be filtered", c)
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	// 13 ranks * 4 suits * 3 decks
	if len(cards) != 13*4*3 {
		t.Errorf("expected %d cards, got %d instead", 13*4*3, len(cards))
	}
}

func TestShuffle(t *testing.T) {
	// make shuffleRand deterministic
	// first call to shuffleRand.Perm(52) should be:
	// [40, 35, ...]
	shuffleRand = rand.New(rand.NewSource(0))

	orig := New()
	first := orig[40]
	second := orig[35]
	_ = second
	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("expected %s, got %s", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("expected %s, got %s", second, cards[1])
	}

}
