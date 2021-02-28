package util

import (
	"path/filepath"
	"testing"
)

func TestIsPathExist(t *testing.T) {
	t.Log(IsPathExist("./file.go"))                       // true
	t.Log(IsPathExist("./file1.go"))                      // false
	t.Log(IsPathExist("../util/"))                        // true
	t.Log(IsPathExist("../util1/"))                       // false
	t.Log(IsPathExist(filepath.Join(HomeDir(), ".ssh")))  // true
	t.Log(IsPathExist(filepath.Join(HomeDir(), ".ssh1"))) // false
}

func TestHomeDir(t *testing.T) {
	t.Log(HomeDir())
}
