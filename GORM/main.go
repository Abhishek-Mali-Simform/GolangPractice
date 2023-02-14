package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type Person struct {
	gorm.Model

	Name  string
	Email string `gorm:"typevarchar(100);unique_index"`
	Books []Book
}

type Book struct {
	gorm.Model

	Title      string
	Author     string
	CallNumber int `gorm:"unique_index"`
	PersonID   int
}

var (
	person = &Person{Name: "Abhishek", Email: "abhishek.m@simformsolutions.com"}
	books  = []Book{
		{Title: "Basics Of Golang", Author: "Google", CallNumber: 1234, PersonID: 1},
		{Title: "Advance Of Golang", Author: "Google", CallNumber: 5678, PersonID: 1},
	}
)

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var db *gorm.DB
var err error

func main() {
	//Loading Environment Variables
	// Check .env File
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORRT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORd")

	//Database Connection String
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s TimeZone=Asia/Shanghai", host, user, dbName, password, dbPort)

	//Opening Connection to database
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully Connected to Database...")
	}

	//Close Connection to database when the main function finishes
	defer func() {
		db, err := db.DB()
		if err != nil {
			fmt.Println(err)
		}
		err = db.Close()
		errorCheck(err)
	}()

	//Make Migrations to the database if they have not already been created
	err = db.AutoMigrate(&Person{})
	errorCheck(err)
	err = db.AutoMigrate(&Book{})
	errorCheck(err)

	//Inserting Data into Database
	//db.Create(person)
	//for index := range books {
	//	db.Create(&books[index])
	//}

	//API rotes
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/person/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/book/{id}", GetBook).Methods("GET")
	router.HandleFunc("/create/person", CreatePerson).Methods("POST")
	router.HandleFunc("/create/book", CreateBook).Methods("POST")
	router.HandleFunc("/update/person/{id}", UpdatePerson).Methods("PUT")
	router.HandleFunc("/update/book/{id}", UpdateBook).Methods("PUT")
	router.HandleFunc("/delete/person/{id}", DeletePerson).Methods("DELETE")
	router.HandleFunc("/delete/book/{id}", DeleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// API Controllers

// People Controllers

func GetPeople(writer http.ResponseWriter, request *http.Request) {
	var people []Person
	var books []Book

	//Retrieve people data from DB
	// SELECT * FROM people
	db.Find(&people)

	for index := range people {
		err = db.Model(&people[index]).Association("Books").Find(&books)
		errorCheck(err)
		people[index].Books = books
	}

	//response on http
	err = json.NewEncoder(writer).Encode(&people)
	errorCheck(err)
}

func GetPerson(writer http.ResponseWriter, request *http.Request) {
	//Get variables from mux
	params := mux.Vars(request)

	var person Person
	var books []Book

	//Only gets first person with similar query in this case id is unique so no problem
	db.First(&person, params["id"])

	//Getting all related books
	err = db.Model(&person).Association("Books").Find(&books)
	errorCheck(err)

	//Adding Found Data To Books
	person.Books = books

	err = json.NewEncoder(writer).Encode(person)
	errorCheck(err)
}

func CreatePerson(writer http.ResponseWriter, request *http.Request) {
	var person Person

	err = json.NewDecoder(request.Body).Decode(&person)
	errorCheck(err)

	createdPerson := db.Create(&person)
	err = createdPerson.Error
	if err != nil {
		err = json.NewEncoder(writer).Encode(err)
		errorCheck(err)
	} else {
		err = json.NewEncoder(writer).Encode(&person)
		errorCheck(err)
	}

}

func UpdatePerson(writer http.ResponseWriter, request *http.Request) {
	//Get variables from mux
	params := mux.Vars(request)

	var person, per Person

	//Getting Updated Data
	err = json.NewDecoder(request.Body).Decode(&per)
	errorCheck(err)

	//Only gets first person with similar query in this case id is unique so no problem
	db.First(&person, params["id"])

	//Adding Found Data To Books
	db.Model(&person).Updates(&per)

	err = json.NewEncoder(writer).Encode(person)
	errorCheck(err)
}

func DeletePerson(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	var person Person

	db.First(&person, params["id"])
	db.Delete(&person)

	err = json.NewEncoder(writer).Encode(&person)
	errorCheck(err)
}

// Book Controllers

func GetBooks(writer http.ResponseWriter, request *http.Request) {
	var books []Book

	db.Find(&books)

	//response on http
	err = json.NewEncoder(writer).Encode(&books)
	errorCheck(err)
}

func GetBook(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var book Book

	db.First(&book, params["id"])

	//response on http
	err = json.NewEncoder(writer).Encode(&book)
	errorCheck(err)
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	var book Book

	err = json.NewDecoder(request.Body).Decode(&book)
	errorCheck(err)

	createdBook := db.Create(&book)
	err = createdBook.Error
	if err != nil {
		err = json.NewEncoder(writer).Encode(err)
		errorCheck(err)
	} else {
		err = json.NewEncoder(writer).Encode(&book)
		errorCheck(err)
	}

}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	//Get variables from mux
	params := mux.Vars(request)

	var book, b Book

	//Getting Updated Data
	err = json.NewDecoder(request.Body).Decode(&b)
	errorCheck(err)

	//Only gets first book with similar query in this case id is unique so no problem
	db.First(&book, params["id"])

	//Adding Found Data To Books
	db.Model(&book).Updates(&b)

	err = json.NewEncoder(writer).Encode(book)
	errorCheck(err)
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	var book Book

	db.First(&book, params["id"])
	db.Delete(&book)

	err = json.NewEncoder(writer).Encode(&book)
	errorCheck(err)
}
