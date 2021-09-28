package dao

import (
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
