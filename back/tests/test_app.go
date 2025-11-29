package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/alexinator1/sumb/back/internal/app"
	"github.com/ilyakaznacheev/cleanenv"
)

func NewApp() *app.App {
	cfg := loadTestConfig()
	testApp, err := app.NewApp(cfg)
	if err != nil {
		panic(fmt.Sprintf("Failed to create test app: %v", err))
	}
	return testApp
}

func loadTestConfig() *app.AppConfig {
	envPath := findNearTestFile()

	fmt.Printf("Loading test config from: %s\n", envPath)

	var cfg app.AppConfig
	if err := cleanenv.ReadConfig(envPath, &cfg); err != nil {
		panic(fmt.Sprintf("Failed to load config from %s: %v", envPath, err))
	}

	return &cfg
}

func findNearTestFile() string {
	_, filename, _, _ := runtime.Caller(0)
	testsDir := filepath.Dir(filename)
	backDir := filepath.Dir(testsDir)
	envPath := filepath.Join(backDir, "test.env")

	if _, err := os.Stat(envPath); err == nil {
		return envPath
	}

	return ""
}
