.PHONY: dev dev-back dev-front down logs init-env clean check-env

# Инициализация .env файла
init-env:
	@if [ ! -f .env ]; then \
		echo "Creating .env from .env.dist..."; \
		cp .env.dist .env; \
		echo "Please edit .env file with your local values"; \
	else \
		echo ".env already exists"; \
	fi

# Проверка .env файла
check-env:
	@if [ ! -f .env ]; then \
		echo "Error: .env file not found. Run 'make init-env' first"; \
		exit 1; \
	fi
	@echo "Environment file check passed!"

# Запуск всей системы
dev: check-env
	docker-compose up

# Запуск в фоне
dev-detached: check-env
	docker-compose up -d

# Только back + БД
dev-back: check-env
	docker-compose up postgres back

# Только front
dev-front: check-env
	docker-compose up front

# Остановка
down:
	docker-compose down

# Остановка с удалением volumes
clean:
	docker-compose down -v

# Просмотр логов
logs:
	docker-compose logs -f

logs-back:
	docker-compose logs -f back

logs-front:
	docker-compose logs -f front

# Пересборка
rebuild: check-env
	docker-compose build --no-cache

rebuild-back: check-env
	docker-compose build back

# Статус сервисов
status:
	docker-compose ps

# БД команды
db-shell: check-env
	docker-compose exec postgres psql -U ${DB_USER:-postgres} -d ${DB_NAME:-pos_dev}