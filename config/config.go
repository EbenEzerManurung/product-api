package config

import (
    "os"
    "time"

    "github.com/joho/godotenv"
)

type Config struct {
    DBHost        string
    DBPort        string
    DBUser        string
    DBPassword    string
    DBName        string
    JWTSecret     string
    AccessExpiry  time.Duration
    RefreshExpiry time.Duration
    ServerPort    string
    CacheTTL      time.Duration
}

func Load() *Config {
    _ = godotenv.Load()

    accessExpiry, _ := time.ParseDuration(getEnv("JWT_ACCESS_EXPIRY", "15m"))
    refreshExpiry, _ := time.ParseDuration(getEnv("JWT_REFRESH_EXPIRY", "168h"))
    cacheTTL, _ := time.ParseDuration(getEnv("CACHE_TTL", "60s"))

    return &Config{
        DBHost:        getEnv("DB_HOST", "127.0.0.1"),
        DBPort:        getEnv("DB_PORT", "3306"),
        DBUser:        getEnv("DB_USER", "root"),
        DBPassword:    getEnv("DB_PASSWORD", ""),
        DBName:        getEnv("DB_NAME", "product_api"),
        JWTSecret:     getEnv("JWT_SECRET", "default-secret-change-me"),
        AccessExpiry:  accessExpiry,
        RefreshExpiry: refreshExpiry,
        ServerPort:    getEnv("SERVER_PORT", "8080"),
        CacheTTL:      cacheTTL,
    }
}

func getEnv(key, fallback string) string {
    if val, ok := os.LookupEnv(key); ok {
        return val
    }
    return fallback
}