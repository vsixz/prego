package orm

import (
	"github.com/vsixz/prego/orm"
	"github.com/vsixz/prego/test/orm/models"
	"testing"
)

type userRepository struct {
	*orm.Repository
}

var this *userRepository

var r *orm.Repository

func init() {
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
	r.Database.AutoMigrate(
		&models.User{},
	)
}

func User() *userRepository {
	if this == nil {
		this = &userRepository{r}
	}

	return this
}

func TestOrm(t *testing.T) {
	//db.AutoMigrate(
	//	&models.User{},
	//)
	// Create
	//for i := 1; i <= 10; i++ {
	//	User().Create(&models.User{Name: fmt.Sprintf("user-%d", i)})
	//}

	// Update
	// update one
	User().Update(&models.User{
		Model: orm.Model{
			Id: "d3a4a6caa2aa41cc94c29b7db77c4074",
		},
		Name: "update-one-test-user",
		Age:  100,
	}, orm.UpdateOption{
		Columns: []string{"name", "age"},
	})
	// update many
	User().Update(&models.User{
		Name: "update-many-test-user",
		Age:  10,
	}, orm.UpdateOption{
		Condition: orm.Condition("name like ? and id != ?", "%user-%", "d3a4a6caa2aa41cc94c29b7db77c4074"),
		Columns:   []string{"name", "age"},
	})

	// List
	//var Users []models.User
	//User().List(&Users, orm.ListOption{
	//	Limit:     20,
	//	OrderBy:   "Name",
	//	Desc:      false,
	//	Condition: orm.Condition("name like ? or name = ?", "%User-110%", "User-101"),
	//})
	//for _, c := range Users {
	//	fmt.Println(c.Name)
	//}

	// Page
	//var Users []models.User
	//total, _ := User().Page(&Users, orm.PageOption{
	//	No:        2,
	//	//Size:      10,
	//	OrderBy:   "created_at",
	//	//Desc:      false,
	//	Condition: orm.Condition("name like ?", "%User-1%"),
	//})
	//fmt.Println(total)
	//for _, c := range Users {
	//	fmt.Println(c.Name)
	//}

	// IsExist & Get & Delete
	//createUser := models.User{
	//	Name: "get-test-user",
	//}
	//User().Create(&createUser)
	//// IsExist
	//fmt.Println(User().IsExist("not-exist-id", &models.User{})) // not exist id
	//fmt.Println(User().IsExist(createUser.Id, &models.User{}))
	//// Get
	//var c models.User
	//User().Get(createUser.Id, &c)
	//fmt.Printf("Name: %s", c.Name)
	//// Delete
	//User().Delete(createUser.Id, &models.User{})
}
