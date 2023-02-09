package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "" +
	"lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]
	//BREAK :  Here try giving two argument it will search for only first argument if found else searches for second as well
	// CONTINUE: It will give actual result with unique value  else without it, it will give value with duplicates
queries:
	for _, q := range query {
	search: // filter method filters and or the hence and the or will never be searched
		for i, w := range words {
			switch q {
			case "and", "or", "the":
				break search
			}
			if strings.EqualFold(q, w) {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				//break queries
				continue queries
			}
		}
	}
}
