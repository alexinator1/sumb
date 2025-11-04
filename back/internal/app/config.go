package app

import (
	"log"

  	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
    DBHost     string `env:"DB_HOST" env-default:"localhost" env-description:"Database host"`
    DBPort     string `env:"DB_PORT" env-default:"5432" env-description:"Database port"`
    DBUser     string `env:"DB_USER" env-default:"postgres" env-description:"Database user"`
    DBPassword string `env:"DB_PASSWORD" env-default:"password" env-description:"Database password"`
    DBName     string `env:"DB_NAME" env-default:"app_db" env-description:"Database name"`
    DBSSLMode  string `env:"DB_SSLMODE" env-default:"disable" env-description:"SSL mode for PostgreSQL"`

    ServerPort string `env:"SERVER_PORT" env-default:"8080"`
    DebugPort  string `env:"DEBUG_PORT" env-default:"4000"`
    GoEnv      string `env:"GO_ENV" env-default:"development"`

    CorsAllowedOrigins string `env:"CORS_ALLOWED_ORIGINS" env-default:"*"`

    JWTSecret      string `env:"JWT_SECRET" env-required:"true" env-description:"JWT signing key"`
    JWTExpireHours int    `env:"JWT_EXPIRE_HOURS" env-default:"24"`
}

func Load() *AppConfig {
	var cfg AppConfig

    // cleanenv сам подхватывает .env если он есть
    // или значения из окружения, если они заданы
    if err := cleanenv.ReadEnv(&cfg); err != nil {
        log.Fatalf("Failed to read environment: %v", err)
    }

    return &cfg
}
