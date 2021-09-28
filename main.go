package main

import (
	"fmt"

	"github.com/AllinChen/systool01/service"
)

func main() {

	// fmt.Println(time.Now().Format("20060102150405"))
	// time.Sleep(time.Second)
	// if fmt.Sprint(time.Now().Clock()) == ""
	// service.SendMail()
	// service.DbBackup()
	sys, _ := service.GetSysRunInfo()
	fmt.Println(*sys)
}
