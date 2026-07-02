package configs

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

func PostgresConfig() *Config {
	return &Config{
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "admin",
		DBPassword: "password",
		DBName:     "inventory-db",
		ServerPort: "8080",
	}
}
