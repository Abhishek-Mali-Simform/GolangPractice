package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var Println = fmt.Println

func main() {
	// file, err := os.Create("data.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// intPrimeList := []int{2, 3, 5, 7, 11}
	// var stringPrimeList []string
	// for _, val := range intPrimeList {
	// 	stringPrimeList = append(stringPrimeList, strconv.Itoa((val)))
	// }
	// for _, num := range stringPrimeList {
	// 	_, err := file.WriteString(num + "\n")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// file, err = os.Open("data.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// scan := bufio.NewScanner(file)
	// for scan.Scan() {
	// 	Println("Prime :", scan.Text())
	// }
	// if err := scan.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	_, err := os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		Println("File Doesn't Exist")
	} else {
		file, err := os.OpenFile("data.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if _, err := file.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}
}
