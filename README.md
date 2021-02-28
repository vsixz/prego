# prego
You may need this golang sdk previous to your development, so... previous-golang => prego.
## Install
```bash
go get github.com/vsixz/prego
```
## conf
**Based on viper(github.com/spf13/viper)**

support json, yaml(yml), ini
```golang

```
## orm
**Use orm no-repos**

```golang
package main
import (
	"github.com/vsixz/prego/log"
	"github.com/vsixz/prego/orm"
	"time"
)

type User struct {
	orm.Model
	Name string `gorm:"size:50" json"name"`
	Age  int    `json:"age"`
}

var r *orm.Repository
func main() {
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
		&User{},
	)
    
    // create user
    r.Create(&User{
		Name: "Jay Chou",
		Age:  time.Now().Year() - 1979,
	})
    
    // first user
    var first User
	r.First(&first, orm.FirstOption{
		Condition: orm.Condition("name = ?", "Jay Chou"),
	})
	log.Debugf("first.Name: %s", first.Name)
    
    // you can refer to other usage in floder ./sample/orm/no-repos
}
```
**Use orm with-repos**

```golang
package main

import (
	"github.com/vsixz/prego/log"
	"github.com/vsixz/prego/orm"
	"time"
)
type User struct {
	orm.Model
	Name string `gorm:"size:50" json"name"`
	Age  int    `json:"age"`
}

type userRepository struct {
	*orm.Repository
}

func user() *userRepository {
	return &userRepository{r}
}

// Override Get method
func (r *userRepository) Get(id string) (User, error) {
	var res User
	err := r.Database.First(&res, "id = ?", id).Error
	return res, err
}

// Add GetByName method
func (r *userRepository) GetByName(name string) (User, error) {
	var res User
	err := r.Database.First(&res, "name = ?", name).Error
	return res, err
}

var r *orm.Repository
func main() {
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
		&User{},
	)
    
    // create user
    user().Create(&User{
		Name: "Jay Chou",
		Age:  time.Now().Year() - 1979,
	})
    
    // first user
    var first User
	user().First(&first, orm.FirstOption{
		Condition: orm.Condition("name = ?", "Jay Chou"),
	})
	log.Debugf("first.Name: %s", first.Name)
    
    // you can refer to other usage in floder ./sample/orm/with-repos
}
```