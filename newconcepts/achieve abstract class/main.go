package main

import "fmt"

type AbstractStruct struct {
	Add func(int) int
}

func main() {
	eg := AbstractStruct{}
	eg.Add = func(i int) int {
		return 10 * i
	}
	fmt.Println(eg.Add(10))
}
