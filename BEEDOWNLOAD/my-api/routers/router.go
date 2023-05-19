// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"my-api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/people",
			beego.NSInclude(
				&controllers.PeopleController{},
			),
		),

		beego.NSNamespace("/books",
			beego.NSInclude(
				&controllers.BooksController{},
			),
		),

		beego.NSNamespace("/book",
			beego.NSInclude(
				&controllers.BookController{},
			),
		),

		beego.NSNamespace("/migrations",
			beego.NSInclude(
				&controllers.MigrationsController{},
			),
		),

		beego.NSNamespace("/somethings",
			beego.NSInclude(
				&controllers.SomethingsController{},
			),
		),

		beego.NSNamespace("/person",
			beego.NSInclude(
				&controllers.PersonController{},
			),
		),

		beego.NSNamespace("/students",
			beego.NSInclude(
				&controllers.StudentsController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
