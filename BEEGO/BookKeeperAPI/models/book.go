package models

type Book struct {
	Id         int     `form:"-"`
	Title      string  `form:"title,text,title:"`
	Author     string  `form:"author,text,author:"`
	CallNumber int     `form:"callNumber,number,callNumber " orm:"unique"`
	Person     *Person `orm:"rel(fk)"`
}
