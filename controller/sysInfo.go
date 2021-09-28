package controller

import (
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func GetSysRunInfo() (error, string) {
	var cmd *exec.Cmd
	cmd = exec.Command("dstat", "-c", "-m", "-n", "-t", "1", "1")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return err, ""
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		return err, ""
	}
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
		return err, ""
	}
	return nil, string(bytes)
}

type Info struct {
	CPU struct {
		Usr int
		Sys int
		Idl int
		Wai int
		Stl int
	}
	MEM struct {
		Used int
		Free int
		Buff int
		Cach int
	}
	NET struct {
		Recv int
		Send int
	}
}

func (SysInfo *Info) GetCPUInfo(infof []string) (err error) {

	infof[3] = strings.Replace(infof[3], "\n", "", -1)
	for ck := 1; ck < 10; ck++ {
		infof[3] = strings.Replace(infof[3], "  ", " ", -1)
	}
	CPUInfof := strings.Split(infof[3], " ")
	for m, ci := range CPUInfof {
		// fmt.Println(m, ci)

		if ci == " " {
			CPUInfof = append(CPUInfof[:m], CPUInfof[m+1:]...)
		}
		if ci == "time" {
			CPUInfof = append(CPUInfof[:m], CPUInfof[m+1:]...)
		}
		// CPUInfof = append(CPUInfof[:m], CPUInfof[m:]...)
	}
	SysInfo.CPU.Usr, err = strconv.Atoi(CPUInfof[1])
	if err != nil {
		return err
	}
	SysInfo.CPU.Sys, err = strconv.Atoi(CPUInfof[2])
	if err != nil {
		return err
	}
	SysInfo.CPU.Idl, err = strconv.Atoi(CPUInfof[3])
	if err != nil {
		return err
	}
	SysInfo.CPU.Wai, err = strconv.Atoi(CPUInfof[4])
	if err != nil {
		return err
	}
	SysInfo.CPU.Stl, err = strconv.Atoi(CPUInfof[5])
	if err != nil {
		return err
	}
	return nil
}

func (SysInfo *Info) GetMEMInfo(infof []string) (err error) {
	infof[4] = strings.Replace(infof[4], "\n", "", -1)
	infof[4] = strings.Replace(infof[4], "M", "", -1)
	for ck := 1; ck < 10; ck++ {
		infof[4] = strings.Replace(infof[4], "  ", " ", -1)
	}
	MEMInfo := strings.Split(infof[4], " ")
	SysInfo.MEM.Used, err = strconv.Atoi(MEMInfo[0])
	if err != nil {
		return err
	}
	SysInfo.MEM.Free, err = strconv.Atoi(MEMInfo[1])
	if err != nil {
		return err
	}
	SysInfo.MEM.Buff, err = strconv.Atoi(MEMInfo[2])
	if err != nil {
		return err
	}
	SysInfo.MEM.Cach, err = strconv.Atoi(MEMInfo[3])
	if err != nil {
		return err
	}
	return nil
}

func (SysInfo *Info) GetNETInfo(infof []string) (err error) {
	for ck := 1; ck < 10; ck++ {
		infof[5] = strings.Replace(infof[5], "  ", " ", -1)
	}
	NETInfof := strings.Split(infof[5], " ")

	SysInfo.NET.Recv, err = strconv.Atoi(NETInfof[1])
	if err != nil {
		return err
	}
	SysInfo.NET.Send, err = strconv.Atoi(NETInfof[2])
	if err != nil {
		return err
	}
	return nil
}
