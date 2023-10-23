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
			l.Info(fmt.Sprintf("%d 端口开放！", port))
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
			l.Info(fmt.Sprintf("%s 主机存活!", ip))
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
	} else if args[0] == "md5" {
		l.Info(fmt.Sprintf("正在计算文件%s的MD5值...", args[1]))
		md5 := scan.Md5EncoderInit(args[1])
		err := md5.Md5Encode()
		if err != nil {
			l.Error(err.Error())
		}
		l.Info("MD5值计算完成!")
		l.Info(md5.GetMd5Code())
	} else if args[0] == "sys" {
		l.Info("正在扫描系统信息...")
		ss := scan.SysScannerInit()
		ss.ScanAll()
		cpuPercent, cpuNumber, memTotal, memUsed, memAva, memPercent, loadInfo, hostname, platform, infoDisk := ss.GetAll()
		l.Info(fmt.Sprintf("CPU使用率: %.3f%%", cpuPercent))
		l.Info(fmt.Sprintf("CPU核心数: %v", cpuNumber))
		l.Info(fmt.Sprintf("总内存: %v GB, 已用内存: %v MB, 可用内存: %v MB, 内存使用率: %.3f %%", memTotal, memUsed, memAva, memPercent))
		l.Info(fmt.Sprintf("系统平均负载: %v", loadInfo))
		l.Info(fmt.Sprintf("hostname is: %v", hostname))
		l.Info(fmt.Sprintf("platform is: %v", platform))
		for _, disk := range infoDisk {
			l.Info(fmt.Sprintf("\n%s", disk))
		}
	}
}
