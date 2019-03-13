package log

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type FileLog struct {
	name       string
	path       string
	normalfile *os.File
	errorfile  *os.File
}

func NewFileLog(name, path string) Log {
	if path == "" {
		path = "."
	}

	if name == "" {
		name = "aglog"
	}

	log := &FileLog{
		name: name,
		path: path,
	}
	log.init()

	return log
}

func (f *FileLog) init() {

	normalName := fmt.Sprintf("%s/%s.log", f.path, f.name)
	normalFile, err := os.OpenFile(normalName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file [ %s ] failed, err: [ %v ] ", normalName, err))
	}
	f.normalfile = normalFile

	errorName := fmt.Sprintf("%s/%s.error.log", f.path, f.name)
	errorFile, err := os.OpenFile(errorName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file [ %s ] failed, err: [ %v ]", errorName, err))
	}
	f.errorfile = errorFile

	return
}

func (f *FileLog) Trace(format string, args ...interface{}) {
	log := generateLogForFile(TraceLevel, format, args...)
	fmt.Fprintln(f.normalfile, log)
}

func (f *FileLog) Debug(format string, args ...interface{}) {
	log := generateLogForFile(DebugLevel, format, args...)
	fmt.Fprintln(f.normalfile, log)
}

func (f *FileLog) Info(format string, args ...interface{}) {
	log := generateLogForFile(InfoLevel, format, args...)
	fmt.Fprintln(f.normalfile, log)
}

func (f *FileLog) Warn(format string, args ...interface{}) {
	log := generateLogForFile(WarnLevel, format, args...)
	fmt.Fprintln(f.errorfile, log)
}

func (f *FileLog) Error(format string, args ...interface{}) {
	log := generateLogForFile(ErrorLevel, format, args...)
	fmt.Fprintln(f.errorfile, log)
}

func (f *FileLog) Fatal(format string, args ...interface{}) {
	log := generateLogForFile(FatalLevel, format, args...)
	fmt.Fprintln(f.errorfile, log)
}

func (f *FileLog) Close() {
	f.normalfile.Close()
	f.errorfile.Close()
	return
}

func generateLogForFile(level, format string, args ...interface{}) (string) {

	now := time.Now().Format("2006-01-02 15:04:05")
	fileName, pkgFunc, lineNum := getLineInfo()
	pkgName := strings.Split(pkgFunc, ".")[0]
	funcName := strings.Split(pkgFunc, ".")[1]
	msg := fmt.Sprintf(format, args...)

	log := fmt.Sprintf(FORMAT, level, now, pkgName, fileName, lineNum, funcName, msg)
	return log
}