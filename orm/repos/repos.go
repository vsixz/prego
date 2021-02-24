package repos

//import (
//	"gorm.io/gorm"
//	"revix/conf"
//	"revix/orm/db"
//	"revix/orm/model"
//)
//
//var (
//	r    repos    = initRepos()
//	User userRepos = r.user()
//)
//
//type repos struct {
//	db *gorm.Database
//}
//
//func initRepos() repos {
//	r := repos{
//		db: db.New(conf.Read().Database.Provider),
//	}
//	r.db.Set("gorm:table_options", "ENGINE=InnoDB charset=utf8").AutoMigrate(
//		&model.User{},
//	)
//	return r
//}
