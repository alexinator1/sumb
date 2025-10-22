package app

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	DBHost     string `env:"DB_HOST"`
	DBPort     string `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
	DBSSLMode  string `env:"DB_SSLMODE"`

	ServerPort string `env:"SERVER_PORT"`
	DebugPort  string `env:"DEBUG_PORT"`
	GoEnv      string `env:"GO_ENV"`

	CorsAllowedOrigins string `env:"CORS_ALLOWED_ORIGINS"`

	JWTSecret      string `env:"JWT_SECRET"`
	JWTExpireHours int    `env:"JWT_EXPIRE_HOURS"`
}

func Load() *AppConfig {
	cfg := &AppConfig{}

	if err := envconfig.Process("", cfg); err != nil {
		log.Fatalf("Failed to load app config: %v", err)
	}

	return cfg
}
