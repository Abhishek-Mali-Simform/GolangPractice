package main

import "fmt"

type Add struct {
	staticVar int
	sum       func(int) func(int) int
}

func adder(valToBeStatic int) func(int) int {
	return func(x int) int {
		valToBeStatic += x
		return valToBeStatic
	}
}

func main() {
	var add Add
	add.sum = adder
	pos, neg := add.sum(add.staticVar), add.sum(add.staticVar)
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	fmt.Println(pos(1), neg(1), "\nCheck Static Variable: ", add.staticVar)
}
