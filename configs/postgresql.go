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
		DBHost:     env("DB_HOST"),
		DBPort:     env("DB_PORT"),
		DBUser:     env("DB_USER"),
		DBPassword: env("DB_PASSWORD"),
		DBName:     env("DB_NAME"),
	}
}

// TODO Подключить viper
func env(key string) string {
	maps := map[string]string{
		"DB_HOST":     "localhost",
		"DB_PORT":     "5432",
		"DB_USER":     "admin",
		"DB_PASSWORD": "password",
		"DB_NAME":     "inventory-db",
	}

	return maps[key]
}
