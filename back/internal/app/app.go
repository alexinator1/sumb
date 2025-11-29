package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// App является основным объектом приложения
type App struct {
	Cfg          *AppConfig
	appProviders *AppProviders
	router       *gin.Engine
}

// NewApp создает новый экземпляр приложения
func NewApp(cfg *AppConfig) (*App, error) {
	// Инициализируем Gin в правильном режиме на основе конфигурации
	if cfg.GoEnv != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	providers, err := NewAppProviders(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create app providers: %w", err)
	}
	if err := providers.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize app providers: %w", err)
	}

	router := buildRouter(providers)

	return &App{
		Cfg:          cfg,
		appProviders: providers,
		router:       router,
	}, nil
}

// Run запускает приложение и все его модули
func (a *App) Run(ctx context.Context) error {
	log.Println("Starting application...")

	// Запускаем HTTP сервер
	serverAddr := fmt.Sprintf(":%s", a.Cfg.ServerPort)
	log.Printf("Starting HTTP server on %s", serverAddr)

	server := &http.Server{
		Addr:    serverAddr,
		Handler: a.router,
	}

	// Запускаем сервер в горутине
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server error: %v", err)
		}
	}()

	return nil
}

// Router возвращает роутер для тестирования
func (a *App) Router() *gin.Engine {
	return a.router
}

func (a *App) DB() *gorm.DB {
	return a.appProviders.DbProvider.DB()
}

// Shutdown gracefully останавливает приложение
func (a *App) Shutdown(ctx context.Context) error {
	log.Println("Shutting down application...")

	// Создаем HTTP сервер с текущим роутером для корректного shutdown
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", a.Cfg.ServerPort),
		Handler: a.router,
	}

	// Gracefully shutdown HTTP сервер
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}

	return nil
}
