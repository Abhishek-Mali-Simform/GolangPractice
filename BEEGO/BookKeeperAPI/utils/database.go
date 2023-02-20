package utils

import (
	"BookKeeperAPI/models"
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/lib/pq"
)

func checkError(msg string, err error) {
	if err != nil {
		logs.Error(msg, " Reason: ", err)
	}
}

func GetDB() {

	var err error
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", EnvConfigs.DBHost, EnvConfigs.DBUser, EnvConfigs.DBName, EnvConfigs.DBPassword, EnvConfigs.DBPort)
	err = orm.RegisterDriver("postgres", orm.DRPostgres)
	checkError("Error Registering Driver.", err)
	err = orm.RegisterDataBase(EnvConfigs.DBAlias, "postgres", dbURI)
	checkError("Error Registering Database.", err)
	orm.RegisterModel(new(models.Person))
	orm.RegisterModel(new(models.Book))
	orm.SetMaxIdleConns(EnvConfigs.DBAlias, 100)
	orm.SetMaxOpenConns(EnvConfigs.DBAlias, 100)
	err = orm.RunSyncdb(EnvConfigs.DBAlias, false, true)
	checkError("Error Syncing Database.", err)

}
