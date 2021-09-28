package service

import "github.com/AllinChen/systool01/dao"

func RecordInfo() error {
	Info, err := GetSysRunInfo()
	if err != nil {
		return err
	}
	dao.InitDB()
	var SI dao.SysInfoImpl
	err = SI.RecordInfo(Info)
	if err != nil {
		return err
	}
	return nil
}
