package main

import (
	"fmt"
	"goscanner/scan"
	"goscanner/utils"
	"strconv"
)

func main() {
	args := utils.ParseArgs()

	if args[0] == "sc" {
		beginport, _ := strconv.Atoi(args[1])
		endport, _ := strconv.Atoi(args[2])
		threads, _ := strconv.Atoi(args[3])
		fmt.Println("正在使用端口扫描器扫描端口...")
		ts := scan.ScannerInit(beginport, endport, threads)
		ts.ScanAll()
		result := ts.GetOpenPorts()
		fmt.Println("扫描完成！")
		for _, port := range result {
			fmt.Printf("%d 端口打开了！\n", port)
		}
	} else if args[0] == "ps" {
		fmt.Println("正在使用Ping扫描器扫描主机...")
		success, _ := strconv.Atoi(args[2])
		ps := scan.PingScannerInit(args[1], success)
		ps.ScanAll()
		result := ps.GetAliveIp()
		fmt.Println("\n扫描完成!")
		for _, ip := range result {
			fmt.Printf("%s 主机存活！\n", ip)
		}
	} else if args[0] == "ms" {
		fmt.Println("正在使用Mysql扫描器扫描主机...")
		ms := scan.MysqlscannerInit(args[1], args[2], args[3])
		ms.ScanAll()
		result := ms.GetMetadata()
		fmt.Println("\n扫描完成!")
		for _, meta := range result {
			fmt.Println(meta)
		}
	}
}
