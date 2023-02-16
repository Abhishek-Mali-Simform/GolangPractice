package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Employee struct {
	gorm.Model
	Name      string
	Languages []Language `json:",omitempty" gorm:"many2many:employee_languages;"`
}

type Language struct {
	gorm.Model
	Name      string
	Employees []Employee `json:",omitempty" gorm:"many2many:employee_languages;"`
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	err error
	db  *gorm.DB
)

func main() {

	//Loading .env file
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("NAME")

	//Connecting With The Database
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	checkError(err)

	//closing on error
	defer func() {
		db, err := db.DB()
		checkError(err)
		err = db.Close()
		checkError(err)
	}()

	//Migrating Struct in Database
	err = db.AutoMigrate(&Employee{})
	checkError(err)
	err = db.AutoMigrate(&Language{})
	checkError(err)

	var (
		emp  Employee
		lang Language
	)

	if check("EMP", &emp) {
		initEmp()
	}
	if check("LANG", &lang) {
		initLang()
	}

	fmt.Println("=======Employee DATA=======")
	findAll("EMP")
	fmt.Println("=======Language DATA=======")
	findAll("LANG")
}

func findAll(name string) {
	switch name {
	case "EMP":
		var emps []Employee
		err := db.Model(&Employee{}).Preload("Languages", func(db *gorm.DB) *gorm.DB {
			return db.Omit("Employees")
		}).Find(&emps).Error
		fmt.Println(err)
		jn, err := json.MarshalIndent(emps, " ", "\t")
		checkError(err)
		fmt.Println(string(jn))
	case "LANG":
		var langs []Language
		err := db.Model(&Language{}).Preload("Employees", func(db *gorm.DB) *gorm.DB {
			return db.Omit("Employees")
		}).Find(&langs).Error
		fmt.Println(err)
		jn, err := json.MarshalIndent(langs, " ", "\t")
		checkError(err)
		fmt.Println(string(jn))
	}
}

func check[T Employee | Language](value string, model *T) bool {
	switch value {
	case "EMP":
		{

			err = db.Where("name = ?", "Abhishek Mali").First(&model).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return true
				} else {
					return false
				}
			}
		}
	case "LANG":
		{
			err = db.Where("name = ?", "English").First(&model).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return true
				} else {
					fmt.Println(err)
					return false
				}
			}
		}
	default:
		return true
	}
	return false
}

func addModel[T Employee | Language](model *T) {
	err = db.Create(&model).Error
	checkError(err)
}

var (
	emp = Employee{
		Name: "Abhishek Mali",
	}
	lang = Language{
		Name: "English",
	}
)

func initEmp() Employee {
	emp.Languages = append(emp.Languages, lang)
	addModel(&emp)
	return emp
}

func initLang() Language {
	lang.Employees = append(lang.Employees, emp)
	addModel(&lang)
	return lang
}
