package entities

import "github.com/vsixz/prego/orm"

type User struct {
	orm.Model
	Name string `gorm:"size:50" json"name"`
	Age  int    `json:"age"`
}
