package main

import (
	"github.com/astaxie/beego"
	_ "gormMigrate/routers"
)

func main() {
	beego.Run()
}
