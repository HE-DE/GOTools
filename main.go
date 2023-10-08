package main

import (
	"fmt"
	"goscanner/tcpscan"
	"goscanner/utils"
	"strconv"
)

func main() {
	args := utils.ParseArgs()

	if args[0] == "tcp" {
		beginport, _ := strconv.Atoi(args[2])
		endport, _ := strconv.Atoi(args[3])
		threads, _ := strconv.Atoi(args[4])
		ts := tcpscan.Init(args[1], beginport, endport, threads)
		ts.ScanAll()
		result := ts.GetOpenPorts()
		for _, port := range result {
			fmt.Printf("\n%d 端口打开了！\n", port)
		}
	}

}
