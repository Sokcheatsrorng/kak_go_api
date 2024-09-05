package configs

import "os"

type Config struct {
    DatabaseDSN string
}

func LoadConfig() *Config {
    return &Config{
        DatabaseDSN: getEnv("DATABASE_DSN", "author_book_api.db"),
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
