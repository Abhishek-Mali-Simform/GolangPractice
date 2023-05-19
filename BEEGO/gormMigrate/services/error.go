package services

import "github.com/astaxie/beego/logs"

func CheckError(msg string, err error) {
	if err != nil {
		logs.Error(msg, " Reason: ", err)
	}
}

func CheckErrorOrSuccess(fail, pass string, err error) {
	if err != nil {
		logs.Error(fail, " Reason: ", err)
	} else {
		logs.Notice(pass)
	}
}
