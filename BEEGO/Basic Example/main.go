package main

import "github.com/beego/beego/v2/server/web"

type MainController struct {
	web.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("Hello World")
}

func main() {
	web.Router("/", &MainController{})
	web.Run()
}
