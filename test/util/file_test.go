package util

import (
	"github.com/vsixz/prego/util"
	"path/filepath"
	"testing"
)

func TestIsPathExist(t *testing.T) {
	t.Log(util.IsPathExist("./file.go"))                            // true
	t.Log(util.IsPathExist("./file1.go"))                           // false
	t.Log(util.IsPathExist("../util/"))                             // true
	t.Log(util.IsPathExist("../util1/"))                            // false
	t.Log(util.IsPathExist(filepath.Join(util.HomeDir(), ".ssh")))  // true
	t.Log(util.IsPathExist(filepath.Join(util.HomeDir(), ".ssh1"))) // false
}

func TestHomeDir(t *testing.T) {
	t.Log(util.HomeDir())
}
