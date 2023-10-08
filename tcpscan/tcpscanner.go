package tcpscan

import (
	"fmt"
	"goscanner/utils"
	"net"
	"sort"
)

type Tcpscanner struct {
	Ip        string
	BeginPort int
	EndPort   int
	Openports []int
	ChanNum   int
}

func Init(ip string, begin int, end int, channum int) *Tcpscanner {
	return &Tcpscanner{ip, begin, end, make([]int, 0), channum}
}

func (t *Tcpscanner) ScanWorker(ports chan int, results chan int) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", t.Ip, port)
		//fmt.Println("正在检查端口：", port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- -1
			continue
		}
		conn.Close()
		results <- port
	}
}

func (t *Tcpscanner) ScanAll() {
	ports := make(chan int, t.ChanNum)
	results := make(chan int)
	bar := utils.CreateBar(t.EndPort - t.BeginPort + 1)
	for i := 0; i < t.ChanNum; i++ {
		go t.ScanWorker(ports, results)
	}

	go func() {
		for i := t.BeginPort; i <= t.EndPort; i++ {
			ports <- i
		}
	}()

	for i := t.BeginPort; i <= t.EndPort; i++ {
		bar.Add(1)
		port := <-results
		if port != -1 {
			t.Openports = append(t.Openports, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(t.Openports)
}

func (t *Tcpscanner) GetOpenPorts() []int {
	return t.Openports
}
