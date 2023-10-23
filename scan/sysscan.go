package scan

import (
	"goscanner/utils"
)

type SysScanner struct {
	//CPU信息
	CpuPercent float64
	CpuNumber  int
	//内存信息
	MemTotal   uint64
	MemUsed    uint64
	MemAva     uint64
	MemPercent float64
	//系统负载
	loadInfo string
	//主机信息
	Hostname string
	Platform string
	//磁盘信息
	InfoDisk []string
}

func SysScannerInit() *SysScanner {
	return &SysScanner{}
}

func (s *SysScanner) ScanAll() {
	s.CpuPercent, s.CpuNumber, _ = utils.GetCpuInfo()
	s.MemTotal, s.MemUsed, s.MemAva, s.MemPercent, _ = utils.GetMemInfo()
	s.loadInfo, _ = utils.GetSysLoad()
	s.Hostname, s.Platform, _ = utils.GetHostInfo()
	s.InfoDisk, _ = utils.GetDiskInfo()
}

func (s *SysScanner) GetAll() (float64, int, uint64, uint64, uint64, float64, string, string, string, []string) {
	return s.CpuPercent, s.CpuNumber, s.MemTotal, s.MemUsed, s.MemAva, s.MemPercent, s.loadInfo, s.Hostname, s.Platform, s.InfoDisk
}
