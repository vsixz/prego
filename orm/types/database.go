package types

type DatabaseConfig struct {
	Driver   DatabaseDriver `json:"driver"`
	Host     string         `json:"host"`
	Port     string         `json:"port"`
	Database string         `json:"database"`
	Username string         `json:"username"`
	Password string         `json:"password"`
}

type DatabaseDriver string

const (
	Mysql      DatabaseDriver = "mysql"
	Sqlserver  DatabaseDriver = "sqlserver"
	Postgresql DatabaseDriver = "postgresql"
	Sqlite     DatabaseDriver = "sqlite"
)
