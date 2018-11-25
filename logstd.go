package aglog

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

func Trace(format string, args ...interface{}) {
	msg := generateLog(format, args...)
	log := fmt.Sprintf(`[%s] %s`, blue(DebugLevel), msg)
	fmt.Println(log)
}

func Debug(format string, args ...interface{}) {
	msg := generateLog(format, args...)
	log := fmt.Sprintf(`[%s] %s`, cyanblue(DebugLevel), msg)
	fmt.Println(log)
}

func Info(format string, args ...interface{}) {
	msg := generateLog(format, args...)
	log := fmt.Sprintf(`[%s] %s`, green(DebugLevel), msg)
	fmt.Println(log)
}

func Warn(format string, args ...interface{}) {
	msg := generateLog(format, args...)
	log := fmt.Sprintf(`[%s] %s`, yellow(DebugLevel), msg)
	fmt.Println(log)
}

func Error(format string, args ...interface{}) {
	msg := generateLog(format, args...)
	log := fmt.Sprintf(`[%s] %s`, red(DebugLevel), msg)
	fmt.Println(log)
}

func Fatal(format string, args ...interface{}) {
	msg := generateLog(format, args...)
	log := fmt.Sprintf(`[%s] %s`, purpose(DebugLevel), msg)
	fmt.Println(log)
}

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

func red(s string) string {
	return fmt.Sprintf("%c[1;40;31m%s%c[0m", 0x1B, s, 0x1B)
}

func green(s string) string {
	return fmt.Sprintf("%c[1;40;32m%s%c[0m", 0x1B, s, 0x1B)
}

func yellow(s string) string {
	return fmt.Sprintf("%c[1;40;33m%s%c[0m", 0x1B, s, 0x1B)
}

func blue(s string) string {
	return fmt.Sprintf("%c[1;40;34m%s%c[0m", 0x1B, s, 0x1B)
}

func purpose(s string) string {
	return fmt.Sprintf("%c[1;40;35m%s%c[0m", 0x1B, s, 0x1B)
}

func cyanblue(s string) string {
	return fmt.Sprintf("%c[1;40;36m%s%c[0m", 0x1B, s, 0x1B)
}
