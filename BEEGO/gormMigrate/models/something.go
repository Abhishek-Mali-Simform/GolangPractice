package models

import (
	"gorm.io/gorm"
	"gormMigrate/services"
)

type Something struct {
	gorm.Model
	Name  string
	Email string
}

func (student *Something) Insert() {
	db := services.NewDatabase()
	err := db.Create(&student).Error
	services.CheckErrorOrSuccess("Error Inserting Data", "Data Inserted Successfully", err)
}

func (student *Something) View() {
	db := services.NewDatabase()
	err := db.First(&student).Error
	services.CheckErrorOrSuccess("Data Not Exists", "Data Found Successfully", err)
}
