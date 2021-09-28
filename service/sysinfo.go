package service

import (
	"strings"

	"github.com/AllinChen/systool01/controller"
)

func GetSysRunInfo() (SysInfo *controller.Info, err error) {
	err, info := controller.GetSysRunInfo()
	if err != nil {
		return nil, err
	}
	SysInfo = new(controller.Info)
	infof := strings.Split(info, "|")
	SysInfo.GetCPUInfo(infof)
	SysInfo.GetMEMInfo(infof)
	SysInfo.GetNETInfo(infof)
	// fmt.Println(SysInfo.CPU, SysInfo.MEM, SysInfo.NET)
	return SysInfo, nil
}

type SysInfo interface {
	GetCPUInfo(infof []string) (err error)
	GetMEMInfo(infof []string) (err error)
	GetNETInfo(infof []string) (err error)
}
