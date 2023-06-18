package main

import (
	"fmt"
	"github.com/fakefloordiv/binsearch"
)

func main() {
	const (
		hello        = "Hello"
		something    = "something"
		disallowed   = "I JUST NEEDED SOME LONG TEST STRINGS"
		someMoreText = "SomeVeryVeryVeryVeryLooooooongKeyThatIsAlsoPossibleInRealProduction"
	)

	m := binsearch.NewContainer[string]()
	m.SetMap(map[string]string{
		hello:        "world",
		something:    "i want",
		disallowed:   "this hashmap is gonna take significant space.",
		someMoreText: "idk what to write about here",
	})

	fmt.Println(m.Get(hello))
	fmt.Println(m.Get(something))
	fmt.Println(m.Get(disallowed))
	fmt.Println(m.Get(someMoreText))
	fmt.Println(m.Get("random text here"))
}
