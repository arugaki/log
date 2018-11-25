package aglog

import (
	"testing"
)

func TestFileLog(t *testing.T) {
	log := NewFileLog("test", ".")
	log.Trace("this is file trace test")
	log.Debug("this is file debug test")
	log.Info("this is file info test")
	log.Warn("this is file warn test")
	log.Error("this is file error test")
	log.Fatal("this is file fatal test")
	log.Close()
}

func TestLog(t *testing.T) {
	Trace("this is file trace test")
	Debug("this is file debug test")
	Info("this is file info test")
	Warn("this is file warn test")
	Error("this is file error test")
	Fatal("this is file fatal test")
}