package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	_, _ = fmt.Scanf("%s\n", &input)
	answer := 1

	for _, ch := range input {
		str := string(ch)
		if str == strings.ToUpper(str) {
			answer++
		}
	}

	fmt.Println(answer)
}
