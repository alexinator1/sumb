package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alexinator1/sumb/back/internal/app"
)

func main() {
	// Загружаем конфигурацию
	cfg := app.Load()

	// Создаем новый экземпляр приложения
	application, err := app.NewApp(cfg)
	if err != nil {
		log.Fatalf("Failed to create application: %v", err)
	}

	// Создаем контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Инициализируем все модули
	// if err := application.Init(ctx); err != nil {
	// 	log.Fatalf("Failed to initialize application: %v", err)
	// }

	// Обработка сигналов для graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Запускаем приложение в горутине
	errChan := make(chan error, 1)
	go func() {
		errChan <- application.Run(ctx)
	}()

	// Ожидаем сигнала завершения или ошибки
	select {
	case err := <-errChan:
		if err != nil {
			log.Printf("Application error: %v", err)
		}
	case sig := <-sigChan:
		log.Printf("Received signal: %v", sig)
	}

	// Graceful shutdown
	if err := application.Shutdown(ctx); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}
}
