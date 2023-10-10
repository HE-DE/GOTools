package utils

import (
	"strconv"

	"github.com/alecthomas/kingpin/v2"
)

var (
	mod       = kingpin.Flag("mod", "指定扫描器").Default("ps").String()
	protocol  = kingpin.Arg("protocol", "指定扫描的协议").Default("tcp").String()
	beginport = kingpin.Arg("beginport", "指定扫描的起始端口").Default("1").Int()
	endport   = kingpin.Arg("endport", "指定扫描的结束端口").Default("65535").Int()
	thread    = kingpin.Arg("thread", "指定扫描的线程数").Default("100").Int()
)

func ParseArgs() []string {
	kingpin.Parse()
	var Args []string
	Args = append(Args, *mod)
	Args = append(Args, *protocol)
	Args = append(Args, strconv.Itoa(*beginport))
	Args = append(Args, strconv.Itoa(*endport))
	Args = append(Args, strconv.Itoa(*thread))
	return Args
}
