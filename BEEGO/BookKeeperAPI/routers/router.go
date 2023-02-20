package routers

import (
	"BookKeeperAPI/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/person/add", &controllers.PersonController{}, "post:Add")
	beego.Router("/person", &controllers.PersonController{}, "get:Person")
	beego.Router("/people", &controllers.PersonController{}, "get:GetPeople")
	beego.Router("/person/edit/:Id([0-9]+)", &controllers.PersonController{}, "get:Edit")
	beego.Router("/person/delete/:Id([0-9]+)", &controllers.PersonController{}, "get:Delete")
	beego.Router("/person/update", &controllers.PersonController{}, "post:Update")
	beego.Router("/book/add", &controllers.BookController{}, "post:Add")
	beego.Router("/book", &controllers.BookController{}, "get:Book")
	beego.Router("/books", &controllers.BookController{}, "get:GetBooks")
	beego.Router("/book/edit/:Id([0-9]+)", &controllers.BookController{}, "get:Edit")
	beego.Router("/book/delete/:Id([0-9]+)", &controllers.BookController{}, "get:Delete")
	beego.Router("/book/update", &controllers.BookController{}, "post:Update")
}
