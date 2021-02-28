package main

import (
	"errors"
	"github.com/vsixz/prego/log"
)

func main() {
	log.Debug("Hello ", "World")
	log.Debugf("Hello, %s", "JAY CHOU")
	log.Info("Hello ", "World")
	log.Infof("Hello, %s", "JAY CHOU")
	log.Warn("Hello ", "World")
	log.Warnf("Hello, %s", "JAY CHOU")
	log.Error("Hello ", "World")
	log.Errorf("Hello, %s", "JAY CHOU")

	log.Error(errors.New("argument is not valid"))
	log.Errorf("Hello, %s", "JAY CHOU")
	log.Health("Hello ", "World")
	log.Healthf("Hello, %s", "JAY CHOU")
	log.Fatal("Hello ", "World")
	log.Fatalf("Hello, %s", "JAY CHOU")
}
