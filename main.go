package main

import (
	"fmt"
	"time"

	"github.com/AllinChen/systool01/service"
)

func main() {

	// fmt.Println(time.Now().Format("20060102150405"))
	// time.Sleep(time.Second)
	// if fmt.Sprint(time.Now().Clock()) == ""

	// service.DbBackup()

	// fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// fmt.Println(service.GetInfo())

	for {
		if service.RecordInfo() == nil {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "  Record Success")
			time.Sleep(10 * time.Second)
		}
		if fmt.Sprint(time.Now().Format("15:04:05")) == "08:00:00" {
			service.SendMail()
		}
	}
	// service.SendMail()

}
