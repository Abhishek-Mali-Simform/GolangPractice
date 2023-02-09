package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type ToDoList struct {
	ToDoCount int
	ToDos     []string
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	errorCheck(err)
}

func getStrings(filename string) []string {
	var lines []string
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return nil
	}
	errorCheck(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	errorCheck(scanner.Err())
	return lines
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Helllo Internet")
}

func interactHandler(writer http.ResponseWriter, request *http.Request) {
	todoVals := getStrings("todos.txt")
	fmt.Printf("%#v\n", todoVals)
	tmpl, err := template.ParseFiles("view.html")
	errorCheck(err)
	todos := ToDoList{
		ToDoCount: len(todoVals),
		ToDos:     todoVals,
	}
	err = tmpl.Execute(writer, todos)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	templ, err := template.ParseFiles("new.html")
	errorCheck(err)
	err = templ.Execute(writer, nil)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	todo := request.FormValue("todo")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("todos.txt", options, os.FileMode(0600))
	errorCheck(err)
	_, err = fmt.Fprintln(file, todo)
	errorCheck(err)
	err = file.Close()
	errorCheck(err)
	http.Redirect(writer, request, "/interact", http.StatusFound)
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/interact", interactHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
