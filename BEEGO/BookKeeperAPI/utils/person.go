package utils

import (
	"BookKeeperAPI/models"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
)

func GetPeople() (people []models.Person) {
	o := orm.NewOrm()
	o.Using(EnvConfigs.DBAlias)
	num, err := o.QueryTable(models.Person{}).All(&people)
	if err != orm.ErrNoRows && num > 0 {
		logs.Debug("Records Found.")
		return people
	} else {
		logs.Debug("Record Not Found. Reason: ", err)
	}
	return nil
}

func GetPerson(Id int) (p models.Person) {
	o := orm.NewOrm()
	o.Using(EnvConfigs.DBAlias)
	p.Id = Id
	err := o.Read(&p, "Id")
	if err != nil {
		logs.Error("Record Not Found. Reason:", err)
	}
	return
}

func UpdatePerson(p *models.Person) (msg string, errs error) {
	o := orm.NewOrm()
	o.Using(EnvConfigs.DBAlias)
	person := GetPerson(p.Id)
	person = *p
	num, err := o.Update(&person)
	msg = string(num) + " Record Updated"
	errs = err
	return
}

func DeletePerson(person *models.Person) (msg string, errs error) {
	o := orm.NewOrm()
	o.Using(EnvConfigs.DBAlias)
	num, err := o.Delete(&models.Person{Id: person.Id})
	errs = err
	msg = string(num) + " Record Deleted."
	return
}
