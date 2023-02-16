package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type User struct {
	gorm.Model
	Name       string `gorm:"index:,unique"`
	CreditCard CreditCard
	DebitCard  DebitCard `gorm:"foreignKey:OwnerID"`
	VisaCard   VisaCard  `gorm:"foreignKey:OwnerName;references:Name"`
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type DebitCard struct {
	gorm.Model
	OwnerID int
	Number  string
}

type VisaCard struct {
	gorm.Model
	Number    string
	OwnerName string
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
	err = db.AutoMigrate(&User{})
	checkError(err)
	err = db.AutoMigrate(&CreditCard{})
	checkError(err)
	err = db.AutoMigrate(&DebitCard{})
	checkError(err)
	err = db.AutoMigrate(&VisaCard{})
	checkError(err)

	var (
		usr    User
		debit  DebitCard
		credit CreditCard
		visa   VisaCard
	)

	if check("USER", &usr) {
		initUser()
	}
	if check("DEBIT", &debit) {
		initDebitCards()
	}
	if check("CREDIT", &credit) {
		initCreditCard()
	}
	if check("VISA", &visa) {
		initVisaCard()
	}
	fmt.Println("=======User DATA=======")
	findAll(&usr)
	fmt.Println("=======Credit Card DATA=======")
	findAll(&credit)
	fmt.Println("=======Debit Card DATA=======")
	findAll(&debit)
	fmt.Println("=======Visa Card DATA=======")
	findAll(&visa)
}

func findAll[T User | CreditCard | DebitCard | VisaCard](model *T) {
	db.Find(&model)
	jn, err := json.MarshalIndent(model, " ", "\t")
	checkError(err)
	fmt.Println(string(jn))
}

func check[T User | CreditCard | DebitCard | VisaCard](value string, model *T) bool {
	switch value {
	case "USER":
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
	case "DEBIT":
		{
			err = db.Where("number = ?", "1607").First(&model).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return true
				} else {
					fmt.Println(err)
					return false
				}
			}
		}
	case "CREDIT":
		{
			err = db.Where("number = ?", "12022022").First(&model).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return true
				} else {
					fmt.Println(err)
					return false
				}
			}
		}
	case "VISA":
		{
			err = db.Where("number = ?", "0609").First(&model).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return true
				} else {
					fmt.Println(err)
					return false
				}
			}
			return false
		}
	default:
		return true
	}
	return false
}

func addModel[T User | DebitCard | VisaCard | CreditCard](model *T) {
	err = db.Create(&model).Error
	checkError(err)
}

func initUser() User {
	user := User{
		Name: "Abhishek Mali",
	}
	addModel(&user)
	return user
}

func initDebitCards() DebitCard {
	debitCard := DebitCard{
		Number:  "1607",
		OwnerID: 1,
	}
	addModel(&debitCard)
	return debitCard
}

func initCreditCard() CreditCard {
	creditCard := CreditCard{
		Number: "12022022",
		UserID: 1,
	}
	addModel(&creditCard)
	return creditCard
}

func initVisaCard() VisaCard {
	visaCard := VisaCard{
		Number:    "0609",
		OwnerName: "Abhishek Mali",
	}
	addModel(&visaCard)
	return visaCard
}
