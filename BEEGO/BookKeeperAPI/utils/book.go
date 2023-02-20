package utils

import (
	"BookKeeperAPI/models"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
)

func GetBooks() (books []models.Book) {
	o := orm.NewOrm()
	o.Using(EnvConfigs.DBAlias)
	num, err := o.QueryTable(models.Book{}).RelatedSel().All(&books)
	if err != orm.ErrNoRows && num > 0 {
		logs.Debug("Records Found.")
		return books
	} else {
		logs.Debug("Record Not Found. Reason: ", err)
	}
	return nil
}

func GetBook(Id int) (book models.Book, person models.Person) {
	o := orm.NewOrm()
	o.Using(EnvConfigs.DBAlias)
	book.Id = Id
	err := o.Read(&book, "Id")
	if err != nil {
		logs.Error("Record Not Found. Reason: ", err)
	}
	person = GetPerson(book.Person.Id)
	return
}

func UpdateBook(b *models.Book) (msg string, errs error) {
	o := orm.NewOrm()
	o.Using(EnvConfigs.DBAlias)
	book, person := GetBook(b.Id)
	book = *b
	book.Person = &person
	num, err := o.Update(&book)
	msg = string(num) + " Records Updated"
	errs = err
	return
}

func DeleteBook(book *models.Book) (msg string, errs error) {
	o := orm.NewOrm()
	o.Using(EnvConfigs.DBAlias)
	num, err := o.Delete(&models.Book{Id: book.Id})
	msg = string(num) + " Record Deleted."
	errs = err
	return
}
