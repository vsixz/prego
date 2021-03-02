package main

import (
	"github.com/vsixz/prego/conf"
	"github.com/vsixz/prego/log"
)

type User struct {
	Id       string
	UserName string `mapstructure:"name"`
}

func main() {
	readFromJsonFile()
	readFromYamlFile()
	readFromYmlFile()
	readFromTomlFile()
	readFromIniFile()
	readFromEnvFile()
	readFromDestFile()
}

func readFromJsonFile() {
	c := conf.Read("./sample/conf/conf.json")
	log.Debug(c.Get("key1"))

	var u User
	c.UnmarshalKey("user", &u)
	log.Debug(u.Id)
	log.Debug(u.UserName)
}

func readFromYamlFile() {
	c := conf.Read("./sample/conf/conf.yaml")
	log.Debug(c.Get("key1"))

	var u User
	c.UnmarshalKey("user", &u)
	log.Debug(u.Id)
	log.Debug(u.UserName)
}

func readFromYmlFile() {
	c := conf.Read("./sample/conf/conf.yml")
	log.Debug(c.Get("key1"))

	var u User
	c.UnmarshalKey("user", &u)
	log.Debug(u.Id)
	log.Debug(u.UserName)
}

func readFromTomlFile() {
	c := conf.Read("./sample/conf/conf.toml")
	log.Debug(c.Get("key1"))

	var u User
	c.UnmarshalKey("user", &u)
	log.Debug(u.Id)
	log.Debug(u.UserName)
}

func readFromIniFile() {
	c := conf.Read("./sample/conf/conf.ini")

	log.Debug(c.Get("default.key1"))
	log.Debug(c.Get("user.id"))
	log.Debug(c.Get("user.name"))
}

func readFromEnvFile() {
	c := conf.Read("./sample/conf/conf.env")

	log.Debug(c.Get("KEY_1"))
	log.Debug(c.Get("USER_ID"))
	log.Debug(c.Get("USER_NAME"))
}

func readFromDestFile() {
	c := conf.Read("./sample/conf/app.conf", "json")
	log.Debug(c.Get("key1"))

	var u User
	c.UnmarshalKey("user", &u)
	log.Debug(u.Id)
	log.Debug(u.UserName)
}
