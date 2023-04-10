package config

type DBType string

const (
	PostgreSQL DBType = "postgres"
	MySQL      DBType = "mysql"
	SQLite3    DBType = "sqlite3"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
	DBType   DBType
}

type Config struct {
	DatabaseConfig
}

func (c *Config) Load() {

}

func Get() *Config {
	return &Config{
		DatabaseConfig{
			Host:     "localhost",
			Port:     "3306",
			Username: "user",
			Password: "123456",
			Database: "account_control_db",
			SSLMode:  "disabled",
			DBType:   "baba",
		},
	}
}
