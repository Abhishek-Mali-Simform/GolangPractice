package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurn = 5
	usage   = "Program will pick %d random Numbers"
	nill    = "Not Number"
	neg     = "Pick Positive Number"
)

var (
	win  string
	lost string
)

func randWin() {
	switch rand.Intn(4) {
	case 0:
		win = "Hurrah You won"
	case 1:
		win = "Yeah you won"
	case 2:
		win = "Won"
	case 3:
		win = "Win"
	default:
		win = "Hii there You won"
	}
}

func randLost() {
	switch rand.Intn(4) {
	case 0:
		lost = "Ughhh You lost"
	case 1:
		lost = "Ohh No you lost"
	case 2:
		lost = "Lost"
	case 3:
		lost = "Didn't Win"
	default:
		lost = "Hii there You lost"
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	l := len(args)
	if l > 1 && l < 2 {
		fmt.Printf(usage, maxTurn)
		return
	}

	guess, err := strconv.Atoi((args[0]))

	if err != nil {
		fmt.Println(nill)
		return
	}

	if guess < 0 {
		fmt.Println(neg)
		return
	}

	guess2, err := strconv.Atoi((args[0]))
	if err != nil {
		fmt.Println(nill)
	}

	if guess2 < 0 {
		fmt.Println(neg)
	}

	for turn := 0; turn < maxTurn; turn++ {
		n := rand.Intn(guess + 1)

		if n == guess || (guess2 != 0 && n == guess) {
			if turn == 0 {
				fmt.Printf("Luckiest Win\t")
			}
			randWin()
			fmt.Println(win)
			return
		}
	}
	randLost()
	fmt.Println(lost)

}
