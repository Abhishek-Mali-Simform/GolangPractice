package models

type Person struct {
	Id    int     `form:"-"`
	Name  string  `form:"name,text,name:" valid:"MinSize(5);MaxSize(20)"`
	Email string  `form:"email,text,email:" orm:"unique"`
	Books []*Book `orm:"reverse(many)"`
}
