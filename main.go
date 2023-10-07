package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports chan int, results chan int) {
	for port := range ports {
		address := fmt.Sprintf("192.168.0.105:%d", port)
		fmt.Println("正在检查端口：", port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- -1
			continue
		}
		conn.Close()
		results <- port
	}
}

func main() {
	ports := make(chan int, 600)
	results := make(chan int)

	var openPorts []int

	for i := 0; i <= cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 65535; i++ {
			ports <- i
		}
	}()

	for i := 1; i <= 65535; i++ {
		port := <-results
		if port != -1 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openPorts)

	for _, port := range openPorts {
		fmt.Println(port, "打开了")
	}
}
