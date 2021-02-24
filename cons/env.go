package cons

var (
	GoEnvironment   = "GO_ENVIRONMENT"
	ConfigFile      = "CONFIG_FILE"
	ConfigMysqlRoot = "CONFIG_MYSQL_ROOT"
)

type RunEnvironment string

var (
	Debugging   RunEnvironment = "Debugging"
	Development RunEnvironment = "Development"
	Testing     RunEnvironment = "Testing"
	Staging     RunEnvironment = "Staging"
	Sandbox     RunEnvironment = "Sandbox"
	Production  RunEnvironment = "Production"
)
