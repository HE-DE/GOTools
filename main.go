package main

import (
	"fmt"
	"goscanner/logger"
	"goscanner/scan"
	"goscanner/utils"
	"strconv"
)

func main() {
	args := utils.ParseArgs()
	l := logger.NewLogger()

	if args[0] == "sc" {
		beginport, _ := strconv.Atoi(args[1])
		endport, _ := strconv.Atoi(args[2])
		threads, _ := strconv.Atoi(args[3])
		l.Info("正在使用端口扫描器扫描端口...")
		ts := scan.ScannerInit(beginport, endport, threads)
		ts.ScanAll()
		result := ts.GetOpenPorts()
		fmt.Println()
		l.Info("扫描完成！")
		for _, port := range result {
			l.Notice(fmt.Sprintf("%d 端口开放！", port))
		}
	} else if args[0] == "ps" {
		l.Info("正在使用Ping扫描器扫描主机...")
		success, _ := strconv.Atoi(args[2])
		ps := scan.PingScannerInit(args[1], success)
		ps.ScanAll()
		result := ps.GetAliveIp()
		fmt.Println()
		l.Info("扫描完成！")
		for _, ip := range result {
			l.Notice(fmt.Sprintf("%s 主机存活!", ip))
		}
	} else if args[0] == "ms" {
		l.Info("正在使用Mysql扫描器扫描主机...")
		ms := scan.MysqlscannerInit(args[1], args[2], args[3])
		ms.ScanAll()
		result := ms.GetMetadata()
		fmt.Println()
		l.Info("扫描完成！")
		for _, meta := range result {
			l.Notice(meta)
		}
	}
}
