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
			log.Fatal(err)
		}
		db.Close()
	}()

	//Make Migrations to the database if they have not already been created
	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Book{})

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
		db.Model(&people[index]).Association("Books").Find(&books)
		people[index].Books = books
	}

	//response on http
	json.NewEncoder(writer).Encode(&people)
}

func GetPerson(writer http.ResponseWriter, request *http.Request) {
	//Get variables from mux
	params := mux.Vars(request)

	var person Person
	var books []Book

	//Only gets first person with similar query in this case id is unique so no problem
	db.First(&person, params["id"])

	//Getting all related books
	db.Model(&person).Association("Books").Find(&books)

	//Adding Found Data To Books
	person.Books = books

	json.NewEncoder(writer).Encode(person)
}

func CreatePerson(writer http.ResponseWriter, request *http.Request) {
	var person Person

	json.NewDecoder(request.Body).Decode(&person)

	createdPerson := db.Create(&person)
	err = createdPerson.Error
	if err != nil {
		json.NewEncoder(writer).Encode(err)
	} else {
		json.NewEncoder(writer).Encode(&person)
	}

}

func UpdatePerson(writer http.ResponseWriter, request *http.Request) {
	//Get variables from mux
	params := mux.Vars(request)

	var person, per Person

	//Getting Updated Data
	json.NewDecoder(request.Body).Decode(&per)

	//Only gets first person with similar query in this case id is unique so no problem
	db.First(&person, params["id"])

	//Adding Found Data To Books
	db.Model(&person).Updates(&per)

	json.NewEncoder(writer).Encode(person)
}

func DeletePerson(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	var person Person

	db.First(&person, params["id"])
	db.Delete(&person)

	json.NewEncoder(writer).Encode(&person)
}

// Book Controllers

func GetBooks(writer http.ResponseWriter, request *http.Request) {
	var books []Book

	db.Find(&books)

	//response on http
	json.NewEncoder(writer).Encode(&books)
}

func GetBook(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var book Book

	db.First(&book, params["id"])

	//response on http
	json.NewEncoder(writer).Encode(&book)
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	var book Book

	json.NewDecoder(request.Body).Decode(&book)

	createdBook := db.Create(&book)
	err = createdBook.Error
	if err != nil {
		json.NewEncoder(writer).Encode(err)
	} else {
		json.NewEncoder(writer).Encode(&book)
	}

}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	var book Book

	db.First(&book, params["id"])
	db.Delete(&book)

	json.NewEncoder(writer).Encode(&book)
}
