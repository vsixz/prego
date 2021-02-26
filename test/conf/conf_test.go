package conf

import (
	"github.com/vsixz/prego/conf"
	"testing"
)

func TestReadFrom(t *testing.T) {
	json := conf.ReadFrom("./conf.json")
	t.Log(json.Get("key1"))

	type User struct {
		Id       string
		UserName string `mapstructure:"name"`
	}
	var u User
	json.UnmarshalKey("user", &u)
	t.Log(u.Id)
	t.Log(u.UserName)
}

func TestReadFromOptions(t *testing.T) {
	json := conf.ReadFromOptions(conf.ConfigOptions{
		Path: "./",
		Type: "json",
		Name: "app.conf",
	})
	t.Log(json.Get("key1"))

	type User struct {
		Id       string
		UserName string `mapstructure:"name"`
	}
	var u User
	json.UnmarshalKey("user", &u)
	t.Log(u.Id)
	t.Log(u.UserName)
}
