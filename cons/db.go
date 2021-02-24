package cons

type DbProvider string

const (
	Mysql      DbProvider = "mysql"
	Sqlserver  DbProvider = "sqlserver"
	Postgresql DbProvider = "postgresql"
	Sqlite     DbProvider = "sqlite"
)
