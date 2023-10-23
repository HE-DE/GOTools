package utils

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/load"
)

func GetCpuInfo() (float64, int, error) {
	cpuPercent, err := cpu.Percent(time.Second, true)
	if err != nil {
		return 0, 0, err
	}
	// fmt.Printf("CPU使用率: %.3f%% \n", cpuPercent[0])
	cpuNumber, err := cpu.Counts(true)
	if err != nil {
		return 0, 0, err
	}
	// fmt.Printf("CPU核心数: %v \n", cpuNumber)
	return cpuPercent[0], cpuNumber, nil
}

func GetMemInfo() (uint64, uint64, uint64, float64, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, 0, 0, err
	}
	// 获取总内存大小，单位GB
	memTotal := memInfo.Total / 1024 / 1024 / 1024
	// 获取已用内存大小，单位MB
	memUsed := memInfo.Used / 1024 / 1024
	// 可用内存大小
	memAva := memInfo.Available / 1024 / 1024
	// 内存可用率
	memUsedPercent := memInfo.UsedPercent
	// fmt.Printf("总内存: %v GB, 已用内存: %v MB, 可用内存: %v MB, 内存使用率: %.3f %% \n", memTotal, memUsed, memAva, memUsedPercent)
	return memTotal, memUsed, memAva, memUsedPercent, nil
}

func GetSysLoad() (string, error) {
	loadInfo, err := load.Avg()
	if err != nil {
		return "", err
	}
	// fmt.Printf("系统平均负载: %v \n", loadInfo)
	return fmt.Sprintf("%v", loadInfo), nil
}

func GetHostInfo() (string, string, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return "", "", err
	}
	// fmt.Printf("hostname is: %v, os platform: %v \n", hostInfo.Hostname, hostInfo.Platform)
	return hostInfo.Hostname, hostInfo.Platform, nil
}

func GetDiskInfo() ([]string, error) {
	diskPart, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}
	resultstorage := make([]string, 0)
	for _, dp := range diskPart {
		temp := fmt.Sprintln(dp)
		diskUsed, _ := disk.Usage(dp.Mountpoint)
		temp += fmt.Sprintf("分区总大小: %d MB \n", diskUsed.Total/1024/1024)
		temp += fmt.Sprintf("分区使用率: %.3f %% \n", diskUsed.UsedPercent)
		temp += fmt.Sprintf("分区inode使用率: %.3f %% \n", diskUsed.InodesUsedPercent)
		resultstorage = append(resultstorage, temp)
	}
	return resultstorage, nil
}
