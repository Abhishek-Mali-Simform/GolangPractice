package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Write Something To mask")
		return
	}

	const (
		link    = "http://"
		lenLink = len(link)
		mask    = '*'
	)

	var (
		text = args[0]
		size = len(text)
		buf  = make([]byte, 0, size)

		in bool
	)

	for i := 0; i < size; i++ {
		if len(text[i:]) > lenLink && text[i:i+lenLink] == link {
			in = true
			buf = append(buf, link...)
			i += lenLink
		}

		c := text[i]

		switch c {
		case ' ', '\t', '\n':
			in = false
		}

		if in {
			c = mask
		}

		buf = append(buf, c)
	}

	fmt.Println(string(buf))
}
