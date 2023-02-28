package conf

var (
	Server *serverConfig
	Zap    *zapConfig
	Mysql  *mysqlConfig
	Redis  *redisConfig
	Email  *emailConfig
)

func Configs() map[string]interface{} {
	return map[string]interface{}{
		"server": &Server,
		"zap":    &Zap,
		"mysql":  &Mysql,
		"redis":  &Redis,
		"email":  &Email,
	}
}

type serverConfig struct {
	ServerName string
	Mode       string
	Addr       string
}

type mysqlConfig struct {
	Dsn       string
	Host      string
	Port      int
	Db        string
	Username  string
	Password  string
	Charset   string
	ParseTime string
	Loc       string
}

type redisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type zapConfig struct {
	DebugFileName string
	InfoFileName  string
	WarnFileName  string
	MaxSize       int
	MaxAge        int
	MaxBackups    int
}

type emailConfig struct {
	Addr     string
	Host     string
	From     string
	To       []string
	Password string
}
