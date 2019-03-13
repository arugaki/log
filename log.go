package log

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type Log interface {
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}

const (
	TraceLevel = "TRACE"
	DebugLevel = "DEBUG"
	InfoLevel  = "INFO"
	WarnLevel  = "WARN"
	ErrorLevel = "ERROR"
	FatalLevel = "FATAL"

	FORMAT = `[%s] time=%s, pkg=%s, file=%s:%d, func=%s, msg=%s`
)

func getLineInfo() (string, string, int) {

	pc, file, line, ok := runtime.Caller(3)
	lineNum := 0
	fileName, funcName := "", ""
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNum = line
	}

	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	return fileName, funcName, lineNum
}

func generateLog(format string, args ...interface{}) (string) {

	now := time.Now().Format("2006-01-02 15:04:05")
	fileName, pkgFunc, lineNum := getLineInfo()
	pkgName := strings.Split(pkgFunc, ".")[0]
	funcName := strings.Split(pkgFunc, ".")[1]
	msg := fmt.Sprintf(format, args...)

	log := fmt.Sprintf(`time=%s, pkg=%s, file=%s:%d, func=%s, msg=%s`,
		now, pkgName, fileName, lineNum, funcName, msg)

	return log
}

