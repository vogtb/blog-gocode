package logger

import "log"

type Logger struct{}

func (l Logger) Info(s string) {
	log.Print("INFO: " + s)
}

func (l Logger) ReturnSame(s string) string {
	return s
}
