package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	store := list{
		{Title: "moby dick", Price: 10, Released: toTimestamp(156446513)},
		{Title: "odyssey", Price: 15, Released: toTimestamp("0326536465")},
		{Title: "hobbit", Price: 16},
	}
	//sort.Sort(store)
	//sort.Sort(sort.Reverse(store))
	//sort.Sort(byReleaseDate(store))
	//sort.Sort(sort.Reverse(byReleaseDate(store)))
	store.discount(.5)
	fmt.Println(store)
	data, err := json.MarshalIndent(store, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(data))

	fmt.Println("\n\n Unmarshaling Data Now\n")

	err = json.Unmarshal([]byte(data), &store)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(store)
}
