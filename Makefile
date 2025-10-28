.PHONY: dev dev-back dev-front down logs init-env clean check-env generate-api generate-back generate-back-main generate-front clean-api help

# Переменные для генерации
OPENAPI_GENERATOR := openapitools/openapi-generator-cli
OPENAPI_SPEC := specs/api/v1/generated/merged-api.yaml
BACK_OUTPUT := back/internal/api/v1/generated
FRONT_OUTPUT := front/src/api/v1/generated

# Подключение модульной генерации
include infra/makefiles/modules.mk

REDOCLY_IMAGE := redocly/redocly-cli:latest

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

# Остановка и удаление контейнеров
down:
	docker-compose down

# Остановка
stop:
	docker-compose stop

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


merge-openapi:
	@echo "🔧 Bundling OpenAPI with Redocly..." 
	@rm -rf ${OPENAPI_SPEC} 
	@docker run --rm -v ${PWD}:/workspace -w /workspace node:18-alpine \
		sh -c "npm install -g @redocly/cli && redocly bundle /workspace/specs/api/v1/openapi.yaml --output $(OPENAPI_SPEC) --ext yaml"
	git add $(OPENAPI_SPEC)
	@echo "✅ OpenAPI bundled and saved to $(OPENAPI_SPEC)"


# Генерация для всего API
generate-api: generate-back generate-front
	@echo "✅ API code generation completed!"


# Генерация только для бекенда (Go) - основной API + все модули
generate-back: generate-back-main generate-modules
	@echo "📝 Adding all generated code to git..."
	@git add $(BACK_OUTPUT)/
	@echo "✅ All generated code added to git"
	@echo "✅ Backend generation completed! (Main API + Modules + Common)"

# Генерация только основного API (без модулей) - теперь только создает структуру папок
generate-back-main:
	@echo "🔧 Preparing generated directory structure..."
	@echo "🧹 Cleaning previous generated code completely..."
	@rm -rf $(BACK_OUTPUT)
	@mkdir -p $(BACK_OUTPUT)
	@echo "✅ Generated directory structure prepared"

# Генерация только для фронтенда javascript
generate-front:
	@echo "🔧 Generating JavaScript client code for Vue 3..."
	@echo "🧹 Cleaning previous generated code..."
	@rm -rf $(FRONT_OUTPUT)
	@mkdir -p $(FRONT_OUTPUT)
	@docker run --rm -v ${PWD}:/local \
		$(OPENAPI_GENERATOR) generate \
		-i /local/$(OPENAPI_SPEC) \
		-g javascript \
		-o /local/$(FRONT_OUTPUT) \
		--additional-properties=\
usePromises=true,\
useES6=true,\
projectName=frontend-api
	@echo "✅ JavaScript client code generated in $(FRONT_OUTPUT)"
	@echo "📝 Adding generated code to git..."
	@git add $(FRONT_OUTPUT)/
	@echo "✅ Generated code added to git"

# Очистка сгенерированного кода
clean-api:
	rm -rf $(BACK_OUTPUT) $(FRONT_OUTPUT)
	@echo "🧹 Generated API code cleaned"

# Валидация OpenAPI спецификации
validate-openapi:
	@echo "🔍 Validating OpenAPI spec..."
	@docker run --rm -v ${PWD}:/local \
		$(OPENAPI_GENERATOR) validate -i /local/$(OPENAPI_SPEC)
	@echo "✅ OpenAPI spec is valid!"

# Просмотр объединенной спецификации
inspect-spec:
	@echo "🔍 Inspecting merged OpenAPI spec..."
	@docker run --rm -v ${PWD}:/local \
		$(OPENAPI_GENERATOR) generate \
		-i /local/$(OPENAPI_SPEC) \
		-g openapi-yaml \
		-o /tmp/openapi-merged \
		--skip-overwrite
	@echo "📄 First 50 lines of merged spec:"
	@cat /tmp/openapi-merged/openapi/openapi.yaml | head -50
	@echo "..."
	@echo "✅ OpenAPI spec successfully merged all refs"

# Генерация Swagger UI документации
swagger-ui:
	@echo "📚 Opening Swagger UI..."
	@docker run --rm -p 8085:8080 \
		-e SWAGGER_JSON=/tmp/openapi.yaml \
		-v ${PWD}/api/openapi.yaml:/tmp/openapi.yaml \
		swaggerapi/swagger-ui

# Полный цикл: валидация + генерация
rebuild-api: validate-openapi clean-api generate-api
	@echo "🎉 API completely rebuilt!"

# Установка зависимостей фронтенда (после генерации)
front-deps:
	@echo "📦 Installing frontend dependencies..."
	cd front && npm install
	@echo "✅ Frontend dependencies installed"

# Запуск бекенда в режиме разработки
# dev-back: generate-back
# 	cd back && go run ./cmd/server

# Запуск фронтенда в режиме разработки  
# dev-front: generate-front front-deps
# 	cd front && npm run dev

# Тесты бекенда
test-back:
	cd back && go test ./...

# Тесты фронтенда
test-front:
	cd front && npm test

# Миграции
migrate-create:
	@read -p "Enter migration name: " name; \
	docker-compose run --rm migrate create -ext sql -dir /migrations -seq $$name

migrate-up:
	docker-compose run --rm migrate up

migrate-down:
	docker-compose run --rm migrate down

# Помощь
help:
	@echo "Available commands:"
	@echo ""
	@echo "Development:"
	@echo "  make dev           - Start full development environment"
	@echo "  make dev-back      - Start only Go backend"
	@echo "  make dev-front     - Start only Vue frontend"
	@echo ""
	@echo "API Generation:"
	@echo "  make generate-api  - Generate code for both backend and frontend"
	@echo "  make generate-back - Generate backend (main API + all modules)"
	@echo "  make generate-back-main - Generate only main API (without modules)"
	@echo "  make generate-front - Generate only TypeScript client code"
	@echo "  make rebuild-api   - Clean, validate and regenerate API"
	@echo "  make validate-openapi - Validate OpenAPI specification"
	@echo "  make inspect-spec  - View merged OpenAPI specification"
	@echo ""
	@echo "Utilities:"
	@echo "  make swagger-ui    - Open Swagger UI for API documentation"
	@echo "  make clean-api     - Remove generated API code"
	@echo "  make front-deps    - Install frontend dependencies"
	@echo "  make test-back     - Run Go tests"
	@echo "  make test-front    - Run Vue tests"
	@echo "  make down          - Stop docker containers"
	@echo "  make logs          - View docker logs"

# По умолчанию показываем помощь
.DEFAULT_GOAL := help
