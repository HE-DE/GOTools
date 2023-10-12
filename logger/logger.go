package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

type LogType int

const (
	DEBUG    = LogType(0)
	INFO     = LogType(1)
	NOTICE   = LogType(2)
	WARNING  = LogType(3)
	ERROR    = LogType(4)
	CRITICAL = LogType(5)
	FATAL    = LogType(6)
)

// 自定义Logger结构类型
type Logger struct {
	mu  sync.Mutex
	out io.Writer

	logFormatFunc func(logType LogType, i interface{}) (string, []interface{}, bool)
	logLevel      LogType
}

// log类型对应的字符串，同一定长处理
var logTypeStrings = func() []string {
	types := []string{"DEBUG", "INFO", "NOTICE", "WARNING", "ERROR", "CRITICAL", "FATAL"}
	maxTypeLen := 0
	for _, v := range types {
		if len(v) > maxTypeLen {
			maxTypeLen = len(v)
		}
	}
	for index, v := range types {
		types[index] = v + strings.Repeat(".", maxTypeLen-len(v))
	}
	return types
}()

// 定义对应类型的颜色
var logTypeColors = []string{"0;35", "1;36", "1;37", "0;33", "1;31", "1;31", "1;31"}

// 创建logger
func NewLogger() *Logger {
	logger := &Logger{}
	logger.Init()
	return logger
}

// 获取Logger的类型
func GetLogTypeString(t LogType) string {
	return logTypeStrings[t]
}

// 初始化Logger
func (l *Logger) Init() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logFormatFunc = l.DefaultLogFormatFunc
	l.out = os.Stdout
	l.logLevel = DEBUG
}

// 设置默认的输出格式
func (l *Logger) DefaultLogFormatFunc(logType LogType, i interface{}) (string, []interface{}, bool) {
	format := "\033[" + logTypeColors[logType] + "m%s [%s] %s \033[0m\n"
	layout := "2006-01-02 15:04:05.999"
	formatTime := time.Now().Format(layout)
	if len(formatTime) != len(layout) {
		// 可能出现结尾是0被省略 如 2006-01-02 15:04:05.9 2006-01-02 15:04:05.99，补上
		formatTime += ".000"[4-(len(layout)-len(formatTime)) : 4]
	}

	values := make([]interface{}, 3)
	values[0] = logTypeStrings[logType]
	values[1] = formatTime
	values[2] = fmt.Sprint(i)

	return format, values, true
}

// 设置日志的级别
func (l *Logger) SetLogLevel(logType LogType) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logLevel = logType
}

// 获取日志的级别
func (l *Logger) GetLogLevel() LogType {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.logLevel
}

// 设置格式化log输出函数
// 函数返回 format 和 对应格式 []interface{}
func (l *Logger) SetLoggerFormat(formatFunc func(logType LogType, i interface{}) (string, []interface{}, bool)) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logFormatFunc = formatFunc
}

// 设置log入口，并进行格式化输出
func (l *Logger) log(logType LogType, i interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.logLevel > logType {
		return
	}

	format, data, isLog := l.logFormatFunc(logType, i)
	if !isLog {
		return
	}

	_, err := fmt.Fprintf(l.out, format, data...)
	if err != nil {
		panic(err)
	}
}

// 不同等级的log输出
func (l *Logger) Debug(i interface{}) {
	l.log(DEBUG, i)
}

func (l *Logger) Info(i interface{}) {
	l.log(INFO, i)
}

func (l *Logger) Notice(i interface{}) {
	l.log(NOTICE, i)
}

func (l *Logger) Warning(i interface{}) {
	l.log(WARNING, i)
}

func (l *Logger) Critical(i interface{}) {
	l.log(CRITICAL, i)
}

func (l *Logger) Error(i interface{}) {
	l.log(ERROR, i)
}

func (l *Logger) Fatal(i interface{}) {
	l.log(FATAL, i)
}
