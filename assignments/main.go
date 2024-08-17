package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	fmt.Println("hello world")

	fmt.Println(arraySign([]int{2, 1}))                    // 1
	fmt.Println(arraySign([]int{-2, 1}))                   // -1
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1})) // 1

	fmt.Println("---")

	fmt.Println(isAnagram("anak", "kana"))       // true
	fmt.Println(isAnagram("anak", "mana"))       // false
	fmt.Println(isAnagram("anagram", "managra")) // true

	fmt.Println("---")

	fmt.Println(findTheDifference("abcd", "abcde")) // 'e'
	fmt.Println(findTheDifference("abcd", "abced")) // 'e'
	fmt.Println(findTheDifference("", "y"))         // 'y'

	fmt.Println("---")

	fmt.Println(canMakeArithmeticProgression([]int{1, 5, 3}))    // true; 1, 3, 5 adalah baris aritmatik +2
	fmt.Println(canMakeArithmeticProgression([]int{5, 1, 9}))    // true; 9, 5, 1 adalah baris aritmatik -4
	fmt.Println(canMakeArithmeticProgression([]int{1, 2, 4, 8})) // false; 1, 2, 4, 8 bukan baris aritmatik, melainkan geometrik x2

	fmt.Println("---")

	tesDeck()
}

// https://leetcode.com/problems/sign-of-the-product-of-an-array
func arraySign(nums []int) int {
	var result int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0
		}
		if nums[i] < 0 {
			result++
		}
	}

	if result%2 != 0 {
		return -1
	}

	return 1
}

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	chars := make(map[rune]int)

	for _, c := range s {
		chars[c]++
	}

	for _, c := range t {
		if number, exists := chars[c]; !exists || number == 0 {
			return false
		}
		chars[c]--
	}

	return len(s) == len(t)
}

// https://leetcode.com/problems/find-the-difference
func findTheDifference(s string, t string) byte {
	var bs, bt byte

	for i := 0; i < len(s); i++ {
		bs += s[i]
	}

	for i := 0; i < len(t); i++ {
		bt += t[i]
	}

	return bt - bs
}

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)

	difference := arr[1] - arr[0]

	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != difference {
			return false
		}
	}

	return true
}

// Deck represent "standard" deck consist of 52 cards
type Deck struct {
	cards []Card
}

// Card represent a card in "standard" deck
type Card struct {
	symbol int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // Ace: 1, Jack: 11, Queen: 12, King: 13
}

// New insert 52 cards into deck d, sorted by symbol & then number.
// [A Spade, 2 Spade,  ..., A Heart, 2 Heart, ..., J Diamond, Q Diamond, K Diamond ]
// assume Ace-Spade on top of deck.
func (d *Deck) New() {
	d.cards = make([]Card, 0, 52)
	for symbol := 0; symbol < 4; symbol++ {
		for number := 1; number <= 13; number++ {
			d.cards = append(d.cards, Card{symbol: symbol, number: number})
		}
	}
}

// PeekTop return n cards from the top
func (d Deck) PeekTop(n int) []Card {
	if n > len(d.cards) {
		n = len(d.cards)
	}
	return d.cards[:n]
}

// PeekTop return n cards from the bottom
func (d Deck) PeekBottom(n int) []Card {
	if n > len(d.cards) {
		n = len(d.cards)
	}
	return d.cards[len(d.cards)-n:]
}

// PeekCardAtIndex return a card at specified index
func (d Deck) PeekCardAtIndex(idx int) Card {
	return d.cards[idx]
}

// Shuffle randomly shuffle the deck
func (d *Deck) Shuffle() {
	rand.New(rand.NewSource(time.Now().UnixNano())).Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

// Cut perform single "Cut" technique. Move n top cards to bottom
// e.g. Deck: [1, 2, 3, 4, 5]. Cut(3) resulting Deck: [4, 5, 1, 2, 3]
func (d *Deck) Cut(n int) {
	if n > 0 && n < len(d.cards) {
		d.cards = append(d.cards[n:], d.cards[:n]...)
	}
}

func (c Card) ToString() string {
	textNum := ""
	switch c.number {
	case 1:
		textNum = "Ace"
	case 11:
		textNum = "Jack"
	case 12:
		textNum = "Queen"
	case 13:
		textNum = "King"
	default:
		textNum = fmt.Sprintf("%d", c.number)
	}
	texts := []string{"Spade", "Heart", "Club", "Diamond"}
	return fmt.Sprintf("%s %s", textNum, texts[c.symbol])
}

func tesDeck() {
	deck := Deck{}
	deck.New()

	top5Cards := deck.PeekTop(5)
	fmt.Println("PeekTop 5")
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}
	fmt.Println("---")
  
  fmt.Println("PeekCardAtIndex index array 10 - 15")
	fmt.Println(deck.PeekCardAtIndex(10).ToString()) // Jack Spade
	fmt.Println(deck.PeekCardAtIndex(11).ToString()) // Queen Spade
	fmt.Println(deck.PeekCardAtIndex(12).ToString()) // King Spade
	fmt.Println(deck.PeekCardAtIndex(13).ToString()) // Ace Heart
	fmt.Println(deck.PeekCardAtIndex(14).ToString()) // 2 Heart
	fmt.Println(deck.PeekCardAtIndex(15).ToString()) // 3 Heart
	fmt.Println("---\n")

	deck.Shuffle()
	top5Cards = deck.PeekTop(5)
	fmt.Println("Deck Shuffle")
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}

	fmt.Println("---")
	deck.New()
	deck.Cut(5)
	bottomCards := deck.PeekBottom(10)
	fmt.Println("Deck Cut")
	for _, c := range bottomCards {
		fmt.Println(c.ToString())
	}
}
