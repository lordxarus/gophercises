package main

import (
	"fmt"
	"strings"
)

func main() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	// fmt.Printf("length: %d\n input: %s\n delta: %d\n", length, input, delta)

	alphaLower := ("abcdefghijklmnopqrstuvwxyz")
	alphaUpper := ("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	ret := ""

	for _, ch := range input {
		switch {
		case strings.ContainsRune(alphaLower, ch):
			ret = ret + string(rotate(ch, delta, alphaLower))
		case strings.ContainsRune(alphaUpper, ch):
			ret = ret + string(rotate(ch, delta, alphaUpper))
		default:
			ret = ret + string(ch)
		}
	}
	fmt.Println(ret)

}

func rotate(s rune, delta int, key string) rune {
	idx := strings.IndexRune(string(key), s)
	if idx < 0 {
		panic("idx < 0")
	}
	idx = (idx + delta) % len(key)
	return []rune(key)[idx]
}
