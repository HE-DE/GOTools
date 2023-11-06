package factory

import (
	"fmt"
	"goscanner/logger"
	"goscanner/scan"
	"strconv"
)

// 策略模式
var l = logger.NewLogger()

type ProcessorStrategy interface {
	Process(args []string)
}

// 实现不同策略的具体结构体
type ScanProcessor struct{}
type PingScanProcessor struct{}
type MysqlScanProcessor struct{}
type FileEncryptProcessor struct{}
type FileDecryptProcessor struct{}
type TarGzipProcessor struct{}
type UnTarGzipProcessor struct{}
type Md5Processor struct{}
type SystemInfoProcessor struct{}
type ImageStitcherProcessor struct{}
type VideoFpsChangeProcessor struct{}

func (p *ScanProcessor) Process(args []string) {
	// 端口扫描逻辑
	beginport, _ := strconv.Atoi(args[0])
	endport, _ := strconv.Atoi(args[1])
	threads, _ := strconv.Atoi(args[2])
	l.Info("正在使用端口扫描器扫描端口...")
	ts := scan.ScannerInit(beginport, endport, threads)
	ts.ScanAll()
	result := ts.GetOpenPorts()
	fmt.Println()
	l.Info("扫描完成！")
	for _, port := range result {
		l.Info(fmt.Sprintf("%d 端口开放！", port))
	}
}

func (p *PingScanProcessor) Process(args []string) {
	// Ping扫描逻辑
	l.Info("正在使用Ping扫描器扫描主机...")
	success, _ := strconv.Atoi(args[1])
	ps := scan.PingScannerInit(args[0], success)
	ps.ScanAll()
	result := ps.GetAliveIp()
	fmt.Println()
	l.Info("扫描完成！")
	for _, ip := range result {
		l.Info(fmt.Sprintf("%s 主机存活!", ip))
	}
}

func (p *MysqlScanProcessor) Process(args []string) {
	// Mysql扫描逻辑
	l.Info("正在使用Mysql扫描器扫描主机数据库...")
	ms := scan.MysqlscannerInit(args[0], args[1], args[2])
	ms.ScanAll()
	result := ms.GetMetadata()
	fmt.Println()
	l.Info("扫描完成！")
	for _, meta := range result {
		l.Notice(meta)
	}
}

func (p *FileEncryptProcessor) Process(args []string) {
	// 文件加密逻辑
	l.Info("正在使用AES加密文件...")
	if len(args[2]) != 16 && len(args[2]) != 24 && len(args[2]) != 32 {
		l.Error("密钥长度错误，请检查！")
		return
	}
	keys := []byte(args[2])
	fe := scan.FileEncoderInit(args[0], keys, args[1])
	err := fe.Encrypt()
	if err != nil {
		l.Error(err.Error())
		return
	}
	l.Info("加密完成！")
}

func (p *FileDecryptProcessor) Process(args []string) {
	// 文件解密逻辑
	l.Info("正在使用AES解密文件...")
	if len(args[2]) != 16 && len(args[2]) != 24 && len(args[2]) != 32 {
		l.Error("密钥长度错误，请检查！")
		return
	}
	keys := []byte(args[2])
	fd := scan.FileDecoderInit(args[0], keys, args[1])
	err := fd.Decrypt()
	if err != nil {
		l.Error(err.Error())
		return
	}
	l.Info("解密完成！")
}

func (p *TarGzipProcessor) Process(args []string) {
	// 文件打包逻辑
	l.Info(fmt.Sprintf("正在打包文件夹%s...", args[0]))
	tar := scan.TarGziperInit(args[0], args[1])
	err := tar.Tar()
	if err != nil {
		l.Error(err.Error())
	}
	l.Info("打包完成！")
}

func (p *UnTarGzipProcessor) Process(args []string) {
	// 文件解包逻辑
	l.Info(fmt.Sprintf("正在解压文件%s...", args[0]))
	tar := scan.UnTarGziperInit(args[0], args[1])
	err := tar.UnTar()
	if err != nil {
		l.Error(err.Error())
	}
	l.Info("解压完成！")
}

func (p *Md5Processor) Process(args []string) {
	// 计算MD5逻辑
	l.Info(fmt.Sprintf("正在计算文件%s的MD5值...", args[0]))
	md5 := scan.Md5EncoderInit(args[0])
	err := md5.Md5Encode()
	if err != nil {
		l.Error(err.Error())
	}
	l.Info("MD5值计算完成!")
	l.Info(md5.GetMd5Code())
}

func (p *SystemInfoProcessor) Process(args []string) {
	// 获取系统信息逻辑
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
	l.Info("系统信息扫描完成！")
}

func (p *ImageStitcherProcessor) Process(args []string) {
	// 图片拼接逻辑
	l.Info("正在拼接图片...")
	maxWidth, _ := strconv.ParseFloat(args[3], 64)
	IS := scan.ImageStitcherInit(args[0], args[1], args[2], maxWidth)
	err := IS.Stitch()
	if err != nil {
		l.Error(err.Error())
	}
	l.Info("图片拼接完成！")
}

func (vc *VideoFpsChangeProcessor) Process(args []string) {
	// 视频帧率转换逻辑
	l.Info("正在转换视频帧率...")
	fps, err := strconv.Atoi(args[2])
	if err != nil {
		l.Error(err.Error())
	}
	VC := scan.InitVideoChanger(args[0], args[1], fps)
	result := VC.Changefps()
	l.Info(result)
	l.Info("视频帧率转换完成！")
}
