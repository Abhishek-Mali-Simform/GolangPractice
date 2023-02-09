package main

import (
	"fmt"
	"time"
)

func main() {

	const (
		width  = 45
		height = 10

		cellEmpty = ' '
		cellBall  = '⚾'

		maxFrames = 1200
		speed     = time.Second / 20
	)

	var (
		px, py int
		vx, vy = 1, 1

		cell rune
	)

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	for i := 0; i < maxFrames; i++ {
		fmt.Print("\033[H\033[2J")
		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}
		board[px][py] = true
		buf := make([]rune, 0, width*height)

		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				//fmt.Print(string(cell), " ") As Print is expensive
				buf = append(buf, cell, ' ')
			}
			//fmt.Println()
			buf = append(buf, '\n')
		}
		fmt.Print(string(buf))
		time.Sleep(speed)
	}

}
