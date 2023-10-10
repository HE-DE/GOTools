package main

import (
	"fmt"
	"goscanner/scan"
	"goscanner/utils"
	"strconv"
)

func main() {
	args := utils.ParseArgs()

	if args[0] == "ps" {
		beginport, _ := strconv.Atoi(args[2])
		endport, _ := strconv.Atoi(args[3])
		threads, _ := strconv.Atoi(args[4])
		fmt.Println("正在使用端口扫描器扫描端口...")
		ts := scan.Init(beginport, endport, threads)
		ts.ScanAll(&args[1])
		result := ts.GetOpenPorts()
		fmt.Println("扫描完成！")
		for _, port := range result {
			fmt.Printf("%d 端口打开了！\n", port)
		}
	}

}
