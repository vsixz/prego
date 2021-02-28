package main

import (
	"fmt"
	"github.com/vsixz/prego/log"
	"github.com/vsixz/prego/orm"
	"github.com/vsixz/prego/sample/orm/with_repos/entities"
	"github.com/vsixz/prego/sample/orm/with_repos/repos"
	"time"
)

func main() {
	CreateIsExistGetFirstDelete()
	List()
	Page()
	All()
}

func CreateIsExistGetFirstDelete() {
	// Create
	create := &entities.User{
		Name: "Jay Chou",
		Age:  time.Now().Year() - 1979,
	}
	repos.User().Create(&create)

	// IsExist
	isExist1 := repos.User().IsExist(create.Id, &entities.User{})
	isExist2 := repos.User().IsExist("not-exist-id", &entities.User{})
	log.Debug("isExist1: %s", isExist1)
	log.Debug("isExist2: %s", isExist2)

	// Get
	get, _ := repos.User().Get(create.Id) // Get method is override.
	log.Debugf("get.Name: %s", get.Name)

	// GetByName
	getByName, _ := repos.User().GetByName("Jay Chou")
	log.Debugf("getByName.Name: %s", getByName.Name)

	// First
	var first entities.User
	repos.User().First(&get, orm.FirstOption{
		Condition: orm.Condition("name = ?", "Jay Chou"),
	})
	log.Debugf("first.Name: %s", first.Name)

	// Delete
	var delete entities.User
	repos.User().Delete(create.Id, &delete)
	log.Debugf("delete.Name: %s", delete.Name)
}

func List() {
	var users []entities.User
	for i := 0; i < 10; i++ {
		users = append(users, entities.User{
			Name: fmt.Sprintf("user-list-%d", i),
			Age:  i,
		})
	}
	repos.User().Database.Create(&users)

	var list []entities.User
	repos.User().List(&list, orm.ListOption{
		Condition: orm.Condition("name like ?", "%user-list-%"),
	})

	log.Debug("list users:")
	for _, u := range list {
		log.Debug(u.Name)
	}

	repos.User().Sweep(&entities.User{}, orm.SweepOption{
		Condition: orm.Condition("name like ?", "%user-list-%"),
	})
}

func Page() {
	var users []entities.User
	for i := 0; i < 30; i++ {
		users = append(users, entities.User{
			Name: fmt.Sprintf("user-page-%d", i),
			Age:  i,
		})
	}
	repos.User().Database.Create(&users)

	var page []entities.User
	repos.User().Page(&page, orm.PageOption{
		Condition: orm.Condition("name like ?", "%user-page-%"),
		No:        2,
		Size:      10,
	})

	log.Debug("page users:")
	for _, u := range page {
		log.Debugf("u.Id: \t%s,\tu.Name: \t%s\t", u.Id, u.Name)
	}

	repos.User().Sweep(&entities.User{}, orm.SweepOption{
		Condition: orm.Condition("name like ?", "%user-page-%"),
	})
}

func All() {
	var users []entities.User
	for i := 0; i < 5; i++ {
		users = append(users, entities.User{
			Name: fmt.Sprintf("user-all-%d", i),
			Age:  i,
		})
	}
	repos.User().Database.Create(&users)

	var all []entities.User
	repos.User().All(&all)
	log.Debug("all users:")
	for _, u := range all {
		log.Debugf("u.Id: \t%s,\tu.Name: \t%s\t", u.Id, u.Name)
	}

	repos.User().Sweep(&entities.User{}, orm.SweepOption{
		Condition: orm.Condition("name like ?", "%user-all-%"),
	})
}
