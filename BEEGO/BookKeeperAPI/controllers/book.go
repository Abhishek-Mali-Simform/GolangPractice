package controllers

import (
	"BookKeeperAPI/models"
	"BookKeeperAPI/utils"
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/sqweek/dialog"
	"strconv"
)

type BookController struct {
	beego.Controller
}

func (bc *BookController) Book() {
	bc.Data["Form"] = &models.Book{}
	bc.Data["People"] = utils.GetPeople()
	bc.TplName = "bookForm.html"
}

func (bc *BookController) Edit() {
	Id, _ := strconv.Atoi(bc.Ctx.Input.Param(":Id"))
	book, _ := utils.GetBook(Id)
	bc.Data["Form"] = &book
	bc.Data["Id"] = Id
	bc.TplName = "editBookForm.html"
}

func (bc *BookController) Add() {
	o := orm.NewOrm()
	o.Using(utils.EnvConfigs.DBAlias)
	pId := bc.GetString("person")
	var p models.Person
	id, err := strconv.Atoi(pId)
	if err != nil {
		logs.Error(err)
	}
	p = utils.GetPerson(id)
	book := models.Book{}
	book.Person = &p
	if err := bc.ParseForm(&book); err != nil {
		logs.Error("Couldn't parse the form. Reason: ", err)
	} else {
		bc.Data["Book"] = book
	}
	if bc.Ctx.Input.Method() == "POST" {
		valid := validation.Validation{}
		isValid, err := valid.Valid(book)
		if err != nil {
			logs.Error(err)
		}
		if !isValid {
			bc.Data["Error"] = valid.ErrorsMap
			logs.Error("Form didn't Validate")
		} else {
			id, err := o.Insert(&book)
			if err != nil {
				logs.Debug("Couldn't insert new book. Reason: ", err)
			} else {
				msg := fmt.Sprintf("Book inserted with id: %d", id)
				logs.Debug(msg)
			}
			bc.Redirect("/books", 302)
		}
	}
}

func (bc *BookController) GetBooks() {
	flash := beego.ReadFromRequest(&bc.Controller)
	if ok := flash.Data["error"]; ok != "" {
		bc.Data["Error"] = ok
	}
	if ok := flash.Data["notice"]; ok != "" {
		bc.Data["notice"] = ok
	}
	bc.Data["Books"] = utils.GetBooks()
	bc.TplName = "viewBooks.html"
}

func (bc *BookController) Update() {
	flash := beego.NewFlash()
	book := models.Book{}
	if err := bc.ParseForm(&book); err != nil {
		logs.Error("Couldn't parse form. Reason: ", err)
	} else {
		bc.Data["Book"] = book
	}
	if bc.Ctx.Input.Method() == "POST" {
		valid := validation.Validation{}
		isValid, err := valid.Valid(book)
		if err != nil {
			logs.Error(err)
		}
		if !isValid {
			bc.Data["Errors"] = valid.ErrorsMap
			logs.Error("Form didn't Validate")
		} else {
			book.Id, err = bc.GetInt("id")
			if err != nil {
				flash.Notice("Record was NOT Updated.")
				flash.Store(&bc.Controller)
				logs.Error("Couldn't convert id from a string to a number.", err)
			}
			msg, err := utils.UpdateBook(&book)
			if err != nil {
				flash.Notice("Record was NOT Updated.")
				flash.Store(&bc.Controller)
				logs.Debug("Couldn't Update new Book. Reason: ", err)
			} else {
				flash.Notice("Record was Updated.")
				flash.Store(&bc.Controller)
				logs.Debug(msg)
			}
			bc.Redirect("/books", 302)
		}
	}

}

func (bc *BookController) Delete() {
	Id, _ := strconv.Atoi(bc.Ctx.Input.Param(":Id"))
	book, _ := utils.GetBook(Id)
	ok := dialog.Message("Do You Want To Delete Data?").Title("Delete Data?").YesNo()
	if ok {
		msg, err := utils.DeleteBook(&book)
		if err == nil {
			logs.Info(msg)
		} else {
			logs.Error("Record couldn't be deleted. Reason: ", err)
		}
	} else {
		logs.Info("You Aborted Deletion.")
	}
	bc.Redirect("/books", 302)
}
