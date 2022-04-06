package common

type MysqlConfig struct {
	DriverName string
	Host       string
	Port       string
	Username   string
	Password   string
	Database   string
	Charset    string
}

type APPConfig struct {
	Addr string
	Port string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	Database string
}

type LogConfig struct {
	Level       string
	Filename    string
	Max_size    int
	Max_age     int
	Max_backups int
}
