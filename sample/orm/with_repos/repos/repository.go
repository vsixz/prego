package repos

import "github.com/vsixz/prego/orm"

var r *orm.Repository

func int() {
	r = &orm.Repository{
		Database: orm.NewDatabase(orm.DatabaseConfig{
			Type:     "mysql",
			Username: "admin",
			Password: "123456",
			Database: "test",
			Host:     "127.0.0.1",
			Port:     3306,
		}),
	}
}
