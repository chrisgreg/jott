package config

type Config struct {
	DB   *DBConfig
	Port int
}

type DBConfig struct {
	Host   string
	User   string
	Pass   string
	DBName string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Host:   "tcp(db:3306)",
			User:   "root",
			Pass:   "root",
			DBName: "jott",
		},
		Port: 3001,
	}
}
