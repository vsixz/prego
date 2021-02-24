package util

import (
	"github.com/vsixz/prego/util"
	"testing"
)

type U struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func TestJsonMarshal(t *testing.T) {
	v := U{
		Name: "jay",
		Age:  41,
	}
	t.Log(util.JsonMarshal(v))
}

func TestJsonUnmarshal(t *testing.T) {
	var u U
	util.JsonUnmarshal(`{"name":"jay","age":41}`, &u)
	t.Log(u.Name)
	t.Log(u.Age)
}

func TestCurrentFuncName(t *testing.T) {
	t.Log(util.CurrentFuncName())
}

func TestCurrentFuncFullName(t *testing.T) {
	t.Log(util.CurrentFuncFullName())
}
