package main

import (
	_ "BookKeeperAPI/routers"
	"BookKeeperAPI/utils"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	utils.InitEnvConfigs()
	utils.GetDB()
	beego.Run()
}
