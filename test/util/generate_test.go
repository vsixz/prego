package util

import (
	"github.com/vsixz/prego/util"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	t.Log(util.GenerateRandomString(16))
	t.Log(util.GenerateRandomString(16))
	t.Log(util.GenerateRandomString(10))
}

func TestGeneratePureRandomString(t *testing.T) {
	t.Log(util.GeneratePureRandomString(10))
	t.Log(util.GeneratePureRandomString(10))
	t.Log(util.GeneratePureRandomString(10))
}

func TestGenUUID(t *testing.T) {
	t.Log(util.GenerateUUID())
	t.Log(util.GeneratePureUUID())
}
