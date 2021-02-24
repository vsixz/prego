package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	Debug("Hello ", "World")
	DebugF("Hello, %s", "JAY CHOU")
	Info("Hello ", "World")
	InfoF("Hello, %s", "JAY CHOU")
	Warn("Hello ", "World")
	WarnF("Hello, %s", "JAY CHOU")
	Error("Hello ", "World")
	ErrorF("Hello, %s", "JAY CHOU")
	Fatal("Hello ", "World")
	FatalF("Hello, %s", "JAY CHOU")
}
