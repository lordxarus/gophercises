package main

import "fmt"

func main() {
	var length, delta int
	var input string
	_, _ = fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	fmt.Printf("length: %d\n input: %s\n delta: %d\n", length, input, delta)
}
