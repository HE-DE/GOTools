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
	} else if args[0] == "fe" {
		l.Info("正在使用AES加密文件...")
		if len(args[3]) != 16 && len(args[3]) != 24 && len(args[3]) != 32 {
			l.Error("密钥长度错误，请检查！")
			return
		}
		keys := []byte(args[3])
		fe := scan.FileEncoderInit(args[1], keys, args[2])
		err := fe.Encrypt()
		if err != nil {
			l.Error(err.Error())
			return
		}
		l.Info("加密完成！")
	} else if args[0] == "fd" {
		l.Info("正在使用AES解密文件...")
		if len(args[3]) != 16 && len(args[3]) != 24 && len(args[3]) != 32 {
			l.Error("密钥长度错误，请检查！")
			return
		}
		keys := []byte(args[3])
		fd := scan.FileDecoderInit(args[1], keys, args[2])
		err := fd.Decrypt()
		if err != nil {
			l.Error(err.Error())
			return
		}
		l.Info("解密完成！")
	} else if args[0] == "tg" {
		l.Info(fmt.Sprintf("正在打包文件夹%s...", args[1]))
		tar := scan.TarGziperInit(args[1], args[2])
		err := tar.Tar()
		if err != nil {
			l.Error(err.Error())
		}
		l.Info("打包完成！")
	} else if args[0] == "ug" {
		l.Info(fmt.Sprintf("正在解压文件%s...", args[1]))
		tar := scan.UnTarGziperInit(args[1], args[2])
		err := tar.UnTar()
		if err != nil {
			l.Error(err.Error())
		}
		l.Info("解压完成！")
	}
}
