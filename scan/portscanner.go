package scan

import (
	"fmt"
	"goscanner/utils"
	"net"
	"sort"
	"time"
)

type Scanner struct {
	Ip        string
	BeginPort int
	EndPort   int
	Openports []int
	ChanNum   int
}

func ScannerInit(begin int, end int, channum int) *Scanner {
	return &Scanner{"127.0.0.1", begin, end, make([]int, 0), channum}
}

func (t *Scanner) ScanWorker(ports chan int, results chan int) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", t.Ip, port)
		//fmt.Println("正在检查端口：", port)
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			results <- -1
			continue
		}
		conn.Close()
		results <- port
	}
}

func (t *Scanner) ScanAll() {
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

func (t *Scanner) GetOpenPorts() []int {
	return t.Openports
}
