package main

import (
	"log"
	"os"
)

type Logger struct {
	debug *log.Logger
	err   *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		debug: log.New(os.Stdout, "[DEBUR] ", 0),
		err:   log.New(os.Stdout, "[ERROR] ", 0),
	}
}

func (l *Logger) Debug(str ...string) {
	l.debug.Println(str)
}

func (l *Logger) Error(err error) {
	l.err.Println(err)
}
