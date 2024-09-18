package main

import (
	"flag"
	"fmt"
	cyoa "gophercises/choose-your-adventure/cyoa"
	"os"
)

func main() {
	path := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("using %s\n", *path)

	f, err := os.Open(*path)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)

}
