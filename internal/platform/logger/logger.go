package logger

import (
	"log"
	"os"
)

var (
	MyLog Logger
)

type Logger struct {
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
}

func (l Logger) Info(v ...interface{}) {
	l.InfoLogger.Println(v...)
}

func (l Logger) Warn(v ...interface{}) {
	l.WarnLogger.Println(v...)
}

func (l Logger) Error(v ...interface{}) {
	l.ErrorLogger.Println(v...)
}

func InitLogger() {
	MyLog = Logger{
		InfoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		WarnLogger:  log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLogger: log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	log.Flags()
}
