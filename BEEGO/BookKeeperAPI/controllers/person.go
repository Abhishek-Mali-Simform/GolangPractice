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

type PersonController struct {
	beego.Controller
}

func (pc *PersonController) Add() {
	o := orm.NewOrm()
	o.Using(utils.EnvConfigs.DBAlias)
	person := models.Person{}
	if err := pc.ParseForm(&person); err != nil {
		logs.Error("Couldn't parse the form. Reason: ", err)
	} else {
		pc.Data["Person"] = person
	}
	if pc.Ctx.Input.Method() == "POST" {
		valid := validation.Validation{}
		isValid, err := valid.Valid(person)
		if err != nil {
			logs.Error(err)
		}
		if !isValid {
			pc.Data["Errors"] = valid.ErrorsMap
			logs.Error("Form didn't Validate")
		} else {
			id, err := o.Insert(&person)
			if err == nil {
				msg := fmt.Sprintf("Person inserted with id: %d", id)
				logs.Debug(msg)
			} else {
				logs.Debug("Couldn't insert new person. Reason: ", err)
			}
			pc.Redirect("/people", 302)
		}
	}
}

func (pc *PersonController) Person() {
	pc.Data["Form"] = &models.Person{}
	pc.TplName = "personForm.html"
}

func (pc *PersonController) GetPeople() {
	flash := beego.ReadFromRequest(&pc.Controller)
	if ok := flash.Data["error"]; ok != "" {
		pc.Data["Error"] = ok
	}
	if ok := flash.Data["notice"]; ok != "" {
		pc.Data["notice"] = ok
	}
	pc.Data["People"] = utils.GetPeople()
	pc.TplName = "viewPeople.html"
}

func (pc *PersonController) Edit() {
	Id, _ := strconv.Atoi(pc.Ctx.Input.Param(":Id"))
	person := utils.GetPerson(Id)
	pc.Data["Form"] = &person
	pc.Data["Id"] = Id
	pc.TplName = "editPersonForm.html"
}

func (pc *PersonController) Update() {
	flash := beego.NewFlash()
	person := models.Person{}
	if err := pc.ParseForm(&person); err != nil {
		logs.Error("Couldn't parse form. Reason: ", err)
	} else {
		pc.Data["Person"] = person
	}
	if pc.Ctx.Input.Method() == "POST" {
		valid := validation.Validation{}
		isValid, err := valid.Valid(person)
		if err != nil {
			logs.Error(err)
		}
		if !isValid {
			pc.Data["Errors"] = valid.ErrorsMap
			logs.Error("Form didn't Validate")
		} else {
			person.Id, err = pc.GetInt("id")
			if err != nil {
				flash.Notice("Record was NOT Updated.")
				flash.Store(&pc.Controller)
				logs.Error("Couldn't convert id from a string to a number.", err)
			}
			msg, err := utils.UpdatePerson(&person)
			if err != nil {
				flash.Notice("Record was NOT Updated.")
				flash.Store(&pc.Controller)
				logs.Debug("Couldn't Update new Person. Reason: ", err)
			} else {
				flash.Notice("Record was Updated.")
				flash.Store(&pc.Controller)
				logs.Debug(msg)
			}
			pc.Redirect("/people", 302)
		}
	}

}

func (pc *PersonController) Delete() {
	Id, _ := strconv.Atoi(pc.Ctx.Input.Param(":Id"))
	person := utils.GetPerson(Id)
	ok := dialog.Message("Do You Want To Delete Data?").Title("Delete Data?").YesNo()
	if ok {
		msg, err := utils.DeletePerson(&person)
		if err == nil {
			logs.Info(msg)
		} else {
			logs.Error("Record couldn't be deleted. Reason: ", err)
		}
	} else {
		logs.Info("You Aborted Deletion.")
	}
	pc.Redirect("/people", 302)
}
