package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// 第一步，创建 2 个日志实例分别用于打印 Info 和 Error 日志
var (
	// [info ] 颜色为蓝色，[error] 为红色 ;使用 log.Lshortfile 支持显示文件名和代码行号
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

var (
	// 暴露 Error，Errorf，Info，Infof 4个方法
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

// 第二步，支持设置日志的层级(InfoLevel, ErrorLevel, Disabled)
const (
	// 三个层级声明为三个常量，通过控制 Output，来控制日志是否打印
	InfoLevel = iota
	ErrorLevel
	Disabled
)

// SetLevel controls log level
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		// ioutil.Discard，即不打印该日志
		errorLog.SetOutput(ioutil.Discard)
	}
	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}
