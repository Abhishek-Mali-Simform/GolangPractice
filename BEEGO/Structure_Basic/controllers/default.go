package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "Hello beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Error"] = "Errors"
	c.TplName = "index.tpl"
}
