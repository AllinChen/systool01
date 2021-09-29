package dao

import (
	"fmt"
	"time"

	"github.com/AllinChen/systool01/controller"
)

type SysInfoImpl struct {
}

var SI SysInfoImpl

func (SysInfoImpl) RecordInfo(SysInfo *controller.Info) error {
	newInfo := ConvertSysType(*SysInfo)
	// fmt.Println(newInfo)
	err := GetDB().Create(newInfo).Error
	return err
}

func (SysInfoImpl) GetOneDayInfo() []Sysinfo {
	var SI []Sysinfo
	d, _ := time.ParseDuration("-24h")
	t1 := time.Now().Add(d).Format("2006-01-02 15:04:05.580337")
	t2 := time.Now().Format("2006-01-02 15:04:05.580337")
	where := fmt.Sprintf("last_modify_date >\"%s\" and last_modify_date <\"%s\"", fmt.Sprint(t1), fmt.Sprint(t2))
	fmt.Println(where)
	GetDB().Find(&SI, where)
	// GetDB().Find(&SI, "id = 3742")
	return SI
}
