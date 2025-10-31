package app

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBProvider struct {
	db *gorm.DB
}

func newDBProvider(cfg *AppConfig) (*DBProvider, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}

	return &DBProvider{
		db: db,
	}, nil
}

// DB returns the database instance
func (p *DBProvider) DB() *gorm.DB {
	return p.db
}
