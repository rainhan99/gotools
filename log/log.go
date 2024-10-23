/*
 * @Author: RainHan
 * @Date: 2024-10-23 11:51:10
 * @LastEditors: RainHan
 * @LastEditTime: 2024-10-23 13:33:40
 * @Description: ===>>>>>>>>>
 * ++
 */

package log

import (
	"log"
	"os"
)

type Logger struct {
	log  *log.Logger
	file string
}

func New(file string) *Logger {
	l := &Logger{
		file: file,
	}
	l.InitLog()
	return l
}

func (l *Logger) InitLog() {
	file, err := os.OpenFile(l.file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		panic(err)
	}
	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.log = logger
}
