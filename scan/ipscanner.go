package scan

import (
	"fmt"
	"goscanner/utils"
	"sync"
	"time"

	"github.com/go-ping/ping"
)

type PingScanner struct {
	Ipaddress string
	Count     int
	AliveIp   []string
}

func PingScannerInit(ip string, count int) *PingScanner {
	return &PingScanner{
		Ipaddress: ip,
		Count:     count,
		AliveIp:   make([]string, 0),
	}
}

var wg sync.WaitGroup

func (ps *PingScanner) Ping(ip string) bool {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		fmt.Printf("创建扫描器失败:%s\n", err.Error())
	}
	// 设置ping包数量
	pinger.Count = ps.Count
	// 设置超时时间
	pinger.Timeout = time.Second * 5
	pinger.Size = 578
	// 设置成特权模式
	pinger.SetPrivileged(true)
	// 运行pinger
	err = pinger.Run()
	if err != nil {
		fmt.Printf("ping异常:%s\n", err.Error())
	}
	stats := pinger.Statistics()
	// 如果回包大于等于1则判为ping通
	if stats.PacketsRecv >= 1 {
		return true
	} else {
		return false
	}
}

func (ps *PingScanner) ScanAll() {
	bar := utils.CreateBar(255)
	for i := 1; i <= 255; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			bar.Add(1)
			ip := fmt.Sprintf("%s.%d", ps.Ipaddress, j)
			if ps.Ping(ip) {
				ps.AliveIp = append(ps.AliveIp, ip)
			}
		}(i)

	}
	wg.Wait()
}

func (ps *PingScanner) ScanNo() {
	bar := utils.CreateBar(255)
	for i := 1; i <= 255; i++ {
		bar.Add(1)
		ip := fmt.Sprintf("%s.%d", ps.Ipaddress, i)
		if ps.Ping(ip) {
			ps.AliveIp = append(ps.AliveIp, ip)
		}
	}
}

func (ps *PingScanner) GetAliveIp() []string {
	return ps.AliveIp
}
