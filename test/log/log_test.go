package log

import (
	"github.com/vsixz/prego/log"
	"testing"
)

func TestLog(t *testing.T) {
	log.Debug("Hello ", "World")
	log.DebugF("Hello, %s", "JAY CHOU")
	log.Info("Hello ", "World")
	log.InfoF("Hello, %s", "JAY CHOU")
	log.Warn("Hello ", "World")
	log.WarnF("Hello, %s", "JAY CHOU")
	log.Error("Hello ", "World")
	log.ErrorF("Hello, %s", "JAY CHOU")
	log.Health("Hello ", "World")
	log.HealthF("Hello, %s", "JAY CHOU")
	log.Fatal("Hello ", "World")
	log.FatalF("Hello, %s", "JAY CHOU")
}
