package log

import (
	"fmt"
	"os"
	"time"
)

type level struct {
	title               string
	foregroundColorCode int
	backgroundColorCode int
}

var (
	debug level = level{
		title:               "DEBUG",
		foregroundColorCode: 34,
		backgroundColorCode: 8,
	}
	info level = level{
		title:               "INFO",
		foregroundColorCode: 32,
		backgroundColorCode: 8,
	}
	warn level = level{
		title:               "WARN",
		foregroundColorCode: 33,
		backgroundColorCode: 8,
	}
	error level = level{
		title:               "ERROR",
		foregroundColorCode: 31,
		backgroundColorCode: 8,
	}
	fatal level = level{
		title:               "FATAL",
		foregroundColorCode: 37,
		backgroundColorCode: 41,
	}
)

func Debug(args ...interface{}) {
	text := fmt.Sprint(args...)
	print(text, debug)
}

func DebugF(format string, args ...interface{}) {
	text := fmt.Sprintf(format, args...)
	print(text, debug)
}

func Info(args ...interface{}) {
	text := fmt.Sprint(args...)
	print(text, info)
}

func InfoF(format string, args ...interface{}) {
	text := fmt.Sprintf(format, args...)
	print(text, info)
}

func Warn(args ...interface{}) {
	text := fmt.Sprint(args...)
	print(text, warn)
}

func WarnF(format string, args ...interface{}) {
	text := fmt.Sprintf(format, args...)
	print(text, warn)
}

func Error(args ...interface{}) {
	text := fmt.Sprint(args...)
	print(text, error)
}

func ErrorF(format string, args ...interface{}) {
	text := fmt.Sprintf(format, args...)
	print(text, error)
}

func Fatal(args ...interface{}) {
	text := fmt.Sprint(args...)
	print(text, fatal)
	os.Exit(1)
}

func FatalF(format string, args ...interface{}) {
	text := fmt.Sprintf(format, args...)
	print(text, fatal)
	os.Exit(1)
}

func print(text string, l level) {
	text = fmt.Sprintf("[%s] [%s] %s", time.Now().Format("2006-01-02 15:04:05.000"), l.title, text)
	fmt.Printf("\033[1;%d;%dm%s\033[0m\n", l.foregroundColorCode, l.backgroundColorCode, text)
}
