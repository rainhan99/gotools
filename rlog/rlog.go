package rlog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

/*
 * @Author: RainHan
 * @Date: 2024-10-23 11:51:10
 * @LastEditors: RainHan
 * @LastEditTime: 2024-10-23 15:09:59
 * @Description: ===>>>>>>>>>
 * ++
 */

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

type Logger struct {
	logLevel LogLevel
	logger   *log.Logger
	file     *os.File
}

func NewLogger(logLevel LogLevel, logFile string) (*Logger, error) {
	var file *os.File
	var err error

	if logFile != "" {
		file, err = os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
		if err != nil {
			return nil, err
		}
	}

	logger := log.New(file, "custom_log", 0)

	return &Logger{
		logLevel: logLevel,
		logger:   logger,
		file:     file,
	}, nil
}

func (l *Logger) Close() {
	if l.file != nil {
		l.file.Close()
	}
}

func (l *Logger) log(level LogLevel, format string, v ...interface{}) {
	if level < l.logLevel {
		return
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	pc, file, line, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	fileName := filepath.Base(file)

	msg := fmt.Sprintf(format, v...)
	logMsg := fmt.Sprintf("[%s] [%s] [%s:%d:%s] %s", now, level.String(), fileName, line, funcName, msg)

	l.logger.Println(logMsg)
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.log(DEBUG, format, v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.log(INFO, format, v...)
}

func (l *Logger) Warning(format string, v ...interface{}) {
	l.log(WARNING, format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.log(ERROR, format, v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	l.log(FATAL, format, v...)
	os.Exit(1)
}
