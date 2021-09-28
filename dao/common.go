package dao

import (
	"github.com/AllinChen/systool01/controller"
)

type Sysinfo struct {
	// ID int64 `gorm:"column:id`

	Cpu_Usr int64 `gorm:"column:cpu_usr`
	Cpu_Sys int64 `gorm:"column:cpu_sys`
	Cpu_Idl int64 `gorm:"column:cpu_idl`
	Cpu_Wai int64 `gorm:"column:cpu_wai`
	Cpu_Stl int64 `gorm:"column:cpu_stl`

	Mem_Used int64 `gorm:"column:mem_used`
	Mem_Free int64 `gorm:"column:mem_free`
	Mem_Buff int64 `gorm:"column:mem_buff`
	Mem_Cach int64 `gorm:"column:mem_cach`

	Net_Recv int64 `gorm:"column:net_recv`
	Net_Send int64 `gorm:"column:net_send`

	// Last_Modify_Date time.Time `gorm:"column:last_modify_date`
}

func ConvertSysType(oldI controller.Info) (newI Sysinfo) {

	newI.Cpu_Usr = int64(oldI.CPU.Usr)
	newI.Cpu_Sys = int64(oldI.CPU.Sys)
	newI.Cpu_Idl = int64(oldI.CPU.Idl)
	newI.Cpu_Wai = int64(oldI.CPU.Wai)
	newI.Cpu_Stl = int64(oldI.CPU.Stl)

	newI.Mem_Used = int64(oldI.MEM.Used)
	newI.Mem_Free = int64(oldI.MEM.Free)
	newI.Mem_Buff = int64(oldI.MEM.Buff)
	newI.Mem_Cach = int64(oldI.MEM.Cach)

	newI.Net_Recv = int64(oldI.NET.Recv)
	newI.Net_Send = int64(oldI.NET.Send)

	return
}
