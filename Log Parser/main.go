package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please write a word")
		return
	}

	query := args[0]
	in := bufio.NewScanner(os.Stdin)

	in.Split(bufio.ScanWords)

	words := make(map[string]bool)
	for in.Scan() {
		word := strings.ToLower(in.Text())

		if len(word) > 2 {
			words[word] = true
		}
	}

	if words[query] {
		fmt.Printf("The input contains %q.\n", query)
		return
	}
	fmt.Printf("The input does not contain %q.\n", query)
}
