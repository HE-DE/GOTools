package utils

import (
	"strconv"

	"github.com/alecthomas/kingpin/v2"
)

var (
	pingscanner = kingpin.Command("ps", "ping扫描网段下的存活IP")
	ip          = pingscanner.Arg("ip", "指定扫描的网段").Required().String()
	count       = pingscanner.Arg("count", "指定回显的包的数量（判断是否存活的标准）").Default("2").Int()

	scanner = kingpin.Command("sc", "扫描端口")
	// protocol  = scanner.Arg("protocol", "指定扫描的协议").Default("tcp").String()
	beginport = scanner.Arg("beginport", "指定扫描的起始端口").Default("1").Int()
	endport   = scanner.Arg("endport", "指定扫描的结束端口").Default("65535").Int()
	thread    = scanner.Arg("thread", "指定扫描的线程数").Default("100").Int()

	mysqlscanner = kingpin.Command("ms", "扫描mysql数据库中表的元信息")
	username     = mysqlscanner.Arg("username", "数据库用户名").Required().String()
	pwd          = mysqlscanner.Arg("pwd", "数据库密码").Required().String()
	dbname       = mysqlscanner.Arg("dbname", "数据库名").Required().String()

	FileEncoder = kingpin.Command("fe", "文件加密")
	FilePath    = FileEncoder.Arg("filepath", "指定要加密的文件路径").Required().String()
	FOutPath    = FileEncoder.Arg("foutpath", "指定加密后的文件路径").Required().String()
	KeyEncoder  = FileEncoder.Arg("key", "指定加密的密钥").Required().String()

	FileDecoder = kingpin.Command("fd", "文件解密")
	FInPath     = FileDecoder.Arg("filepath", "指定要解密的文件路径").Required().String()
	FileOutPath = FileDecoder.Arg("foutpath", "指定解密后的文件路径").Required().String()
	KeyDecoder  = FileDecoder.Arg("key", "指定解密的密钥").Required().String()
)

func ParseArgs() []string {
	var Args []string
	switch kingpin.Parse() {
	case "ps":
		Args = append(Args, "ps")
		Args = append(Args, *ip)
		Args = append(Args, strconv.Itoa(*count))
		return Args
	case "sc":
		Args = append(Args, "sc")
		// Args = append(Args, *protocol)
		Args = append(Args, strconv.Itoa(*beginport))
		Args = append(Args, strconv.Itoa(*endport))
		Args = append(Args, strconv.Itoa(*thread))
		return Args
	case "ms":
		Args = append(Args, "ms")
		Args = append(Args, *username)
		Args = append(Args, *pwd)
		Args = append(Args, *dbname)
		return Args
	case "fe":
		Args = append(Args, "fe")
		Args = append(Args, *FilePath)
		Args = append(Args, *FOutPath)
		Args = append(Args, *KeyEncoder)
		return Args
	case "fd":
		Args = append(Args, "fd")
		Args = append(Args, *FInPath)
		Args = append(Args, *FileOutPath)
		Args = append(Args, *KeyDecoder)
		return Args
	default:
		return nil
	}
}
