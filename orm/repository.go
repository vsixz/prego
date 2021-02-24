package orm

type Repository struct {
	Database *Database
}
//
//var Repos Repository
//
//func init() {
//	configPath := os.Getenv(cons.ConfigFile)
//
//	if len(configPath) == 0 {
//		configPath = "./conf/conf.yaml"
//		//panic("can not read config")
//	}
//	c := conf.ReadFrom(configPath)
//	var dbConf types.DatabaseConfig
//	c.UnmarshalKey("db", &dbConf)
//	//conf.ReadFrom(configPath).Get("db").Unmarshal(&c)
//	//dbConf:=c.(types.DatabaseConfig)
//	Repos = Repository{
//		Database: New(dbConf),
//	}
//}

type S struct {
	Runway string `json:"runway"`
}

func (r Repository) Create(entry interface{})error{
	return r.Database.Create(&entry).Error
}

func (r Repository) Get(id interface{}, entry interface{}) error {
	if err := r.Database.First(&entry, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r Repository) Delete(id interface{}, entry interface{}) error {
	return r.Database.Where("id = ?", id).Delete(&entry).Error
}
