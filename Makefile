.PHONY: dev dev-back dev-back-down dev-front dev-back-test dev-back-test-down down logs init-env clean check-env generate-api generate-back generate-back-main generate-front clean-api help init-env-test check-env-test integration-back-tests unit-back-tests test-back test-front

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
		if [ -f .env.dist ]; then \
			echo "Creating .env from .env.dist..."; \
			cp .env.dist .env; \
			echo "Please edit .env file with your local values"; \
		else \
			echo "❗ .env.dist not found. Creating default .env..."; \
			echo "GO_ENV=development"                                  > .env; \
			echo "SERVER_PORT=8080"                                   >> .env; \
			echo "DEBUG_PORT=40000"                                   >> .env; \
			echo "BACKEND_PORT=8080"                                  >> .env; \
			echo "BACKEND_DEBUG_PORT=40000"                            >> .env; \
			echo "DB_HOST=localhost"                                  >> .env; \
			echo "DB_PORT=5432"                                       >> .env; \
			echo "DB_USER=postgres"                                   >> .env; \
			echo "DB_PASSWORD=password"                               >> .env; \
			echo "DB_NAME=sumb"                                       >> .env; \
			echo "DB_SSLMODE=disable"                                 >> .env; \
			echo "JWT_SECRET=dev-secret"                              >> .env; \
			echo "JWT_EXPIRE_HOURS=72"                                >> .env; \
			echo "CORS_ALLOWED_ORIGINS=*"                             >> .env; \
			echo "✅ Default .env created. Review and adjust if needed."; \
		fi; \
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
dev-a: check-env
	docker-compose up
# Запуск в фоне
dev: check-env
	docker-compose up -d

# Только back + БД (back запускается нативно для дебаггера)
dev-back: check-env
	@echo "🚀 Starting dependencies (postgres, migrate) in Docker..."
	@docker-compose -f docker-compose.deps.yml --env-file .env up -d postgres
	@echo "⏳ Waiting for postgres to be ready..."
	@timeout=30; \
	while [ $$timeout -gt 0 ]; do \
		if docker-compose -f docker-compose.deps.yml --env-file .env exec -T postgres pg_isready -U $${DB_USER:-postgres} > /dev/null 2>&1; then \
			break; \
		fi; \
		echo "   Waiting... ($$timeout)"; \
		sleep 1; \
		timeout=$$((timeout - 1)); \
	done; \
	if [ $$timeout -eq 0 ]; then \
		echo "❌ Timeout waiting for postgres"; \
		exit 1; \
	fi
	@echo "✅ postgres is ready"
	@echo "🔄 Running migrations..."
	@docker-compose -f docker-compose.deps.yml --env-file .env run --rm migrate || true
	@echo "✅ Dependencies are ready!"
	@echo ""
	@echo "📝 To run the Go app natively:"
	@echo "   1. Set environment variables from .env file"
	@echo "   2. Run: cd back && go run ./cmd/server"
	@echo "   3. Or use VS Code debugger (F5)"
	@echo ""
	@echo "📊 Dependencies status:"
	@docker-compose -f docker-compose.deps.yml --env-file .env ps
	@echo ""
	@echo "⚠️  To stop dependencies: make dev-back-down"

# Остановка зависимостей для dev-back
dev-back-down:
	@echo "🛑 Stopping dependencies..."
	@docker-compose -f docker-compose.deps.yml --env-file .env down

# Тестовое окружение для локальной отладки тестов
dev-back-test: check-env-test
	@echo "🚀 Starting test dependencies (postgres-test) in Docker..."
	@docker-compose -f docker-compose.test.yml --env-file back/test.env up -d postgres-test
	@echo "⏳ Waiting for postgres-test to be ready..."
	@timeout=30; \
	while [ $$timeout -gt 0 ]; do \
		if docker-compose -f docker-compose.test.yml --env-file back/test.env exec -T postgres-test pg_isready -U postgres > /dev/null 2>&1; then \
			break; \
		fi; \
		echo "   Waiting... ($$timeout)"; \
		sleep 1; \
		timeout=$$((timeout - 1)); \
	done; \
	if [ $$timeout -eq 0 ]; then \
		echo "❌ Timeout waiting for postgres-test"; \
		exit 1; \
	fi
	@echo "✅ postgres-test is ready"
	@echo "🔄 Running migrations..."
	@docker-compose -f docker-compose.test.yml --env-file back/test.env run --rm migrate || true
	@echo "✅ Test dependencies are ready!"
	@echo ""
	@echo "🧪 To run tests locally with this DB:"
	@echo "   1. Export variables from test.env (e.g. 'set -o allexport && source back/test.env && set +o allexport')"
	@echo "   2. Run: cd back && go test ./..."
	@echo "   3. For integration focus: go test ./... -tags=integration"
	@echo ""
	@echo "📊 Test dependencies status:"
	@docker-compose -f docker-compose.test.yml --env-file back/test.env ps
	@echo ""
	@echo "⚠️  To stop test dependencies: make dev-back-test-down"

dev-back-test-down:
	@echo "🛑 Stopping test dependencies..."
	@docker-compose -f docker-compose.test.yml --env-file back/test.env down

# Остановка и удаление контейнеров
down:
	docker-compose down -v

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

# Запуск back в Docker с Delve (для удаленной отладки)
back-debug: check-env
	@echo "🚀 Starting Postgres dependency..."
	@docker-compose -f docker-compose.deps.yml --env-file .env up -d postgres
	@echo "⏳ Waiting for postgres to be ready..."
	@timeout=30; \
	while [ $$timeout -gt 0 ]; do \
		if docker-compose -f docker-compose.deps.yml --env-file .env exec -T postgres pg_isready -U $${DB_USER:-postgres} > /dev/null 2>&1; then \
			break; \
		fi; \
		echo "   Waiting... ($$timeout)"; \
		sleep 1; \
		timeout=$$((timeout - 1)); \
	done; \
	if [ $$timeout -eq 0 ]; then \
		echo "❌ Timeout waiting for postgres"; \
		exit 1; \
	fi
	@echo "✅ postgres is ready"
	@echo "🔄 Running migrations (non-fatal)..."
	@docker-compose -f docker-compose.deps.yml --env-file .env run --rm migrate || echo "⚠️  Migrations failed or are not needed; continuing for debug"
	@echo "🧰 Building and running backend with Delve in Docker..."
	@echo "   Debug port: $${BACKEND_DEBUG_PORT:-40000} (container and host)"
	@docker-compose --env-file .env run --rm --no-deps --service-ports back \
		sh -lc 'go install github.com/go-delve/delve/cmd/dlv@latest \
		&& dlv debug ./cmd/server --headless --listen=:${DEBUG_PORT:-40000} --api-version=2 --accept-multiclient'
	@echo "✅ Backend stopped"
	@echo "ℹ️  To connect: Delve (localhost:$${BACKEND_DEBUG_PORT:-40000})"

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

# Генерация для всего API
generate-api: generate-back generate-front
	@echo "✅ API code generation completed!"


# Генерация только для бекенда (Go) - основной API + все модули
generate-back: generate-modules


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

# Инициализация test.env файла из test.env.dist
init-env-test:
	@if [ ! -f back/test.env ]; then \
		if [ -f back/test.env.dist ]; then \
			echo "Creating test.env from test.env.dist..."; \
			cp back/test.env.dist back/test.env; \
			echo "✅ back/test.env file created"; \
		else \
			echo "❌ Error: test.env.dist not found. Creating default test.env..."; \
			echo "# Test Environment Configuration" > test.env; \
			echo "DB_HOST=postgres-test" >> test.env; \
			echo "DB_PORT=5432" >> test.env; \
			echo "DB_USER=postgres" >> test.env; \
			echo "DB_PASSWORD=test_password" >> test.env; \
			echo "DB_NAME=sumb_test" >> test.env; \
			echo "DB_SSLMODE=disable" >> test.env; \
			echo "BACKEND_PORT=8081" >> test.env; \
			echo "SERVER_PORT=8081" >> test.env; \
			echo "GO_ENV=test" >> test.env; \
			echo "JWT_SECRET=test-secret-key" >> test.env; \
			echo "✅ Default test.env file created"; \
		fi; \
	else \
		echo "✅ test.env already exists"; \
	fi

# Проверка test.env файла
check-env-test:
	@if [ ! -f back/test.env ]; then \
		echo "Error: test.env file not found. Run 'make init-env-test' first"; \
		exit 1; \
	fi
	@echo "✅ Test environment file check passed!"

# Интеграционные тесты бекенда в Docker контейнере
integration-back-tests: check-env-test
	@echo "🧪 Running backend integration tests in Docker container..."
	@echo "📦 Starting test services (postgres-test)..."
	@docker-compose -f docker-compose.test.yml --env-file test.env up -d postgres-test
	@echo "⏳ Waiting for postgres-test to be ready..."
	@timeout=30; \
	while [ $$timeout -gt 0 ]; do \
		if docker-compose -f docker-compose.test.yml --env-file test.env exec -T postgres-test pg_isready -U postgres > /dev/null 2>&1; then \
			break; \
		fi; \
		echo "   Waiting... ($$timeout)"; \
		sleep 1; \
		timeout=$$((timeout - 1)); \
	done; \
	if [ $$timeout -eq 0 ]; then \
		echo "❌ Timeout waiting for postgres-test"; \
		docker-compose -f docker-compose.test.yml --env-file test.env down; \
		exit 1; \
	fi
	@echo "✅ postgres-test is ready"
	@echo "🔄 Running migrations..."
	@docker-compose -f docker-compose.test.yml --env-file test.env run --rm migrate || true
	@echo "🧪 Running Go integration tests..."
	@docker-compose -f docker-compose.test.yml --env-file test.env run --rm \
		-e CGO_ENABLED=1 \
		-e DB_HOST=postgres-test \
		-e DB_PORT=5432 \
		-e DB_USER=postgres \
		-e DB_PASSWORD=test_password \
		-e DB_NAME=sumb_test \
		-e DB_SSLMODE=disable \
		-e SERVER_PORT=8081 \
		back-test \
		sh -c "cd /app && go test -v -tags=integration ./internal/domain/business/api/v1/handler/... -run 'Test.*Integration'" || \
		(echo "❌ Integration tests failed" && docker-compose -f docker-compose.test.yml --env-file test.env down && exit 1)
	@echo "🧹 Cleaning up test services..."
	@docker-compose -f docker-compose.test.yml --env-file test.env down
	@echo "✅ Integration tests completed!"

# Unit тесты бекенда в Docker контейнере
unit-back-tests: check-env-test
	@echo "🧪 Running backend unit tests in Docker container..."
	@echo "📦 Starting test services (postgres-test)..."
	@docker-compose -f docker-compose.test.yml --env-file test.env up -d postgres-test
	@echo "⏳ Waiting for postgres-test to be ready..."
	@timeout=30; \
	while [ $$timeout -gt 0 ]; do \
		if docker-compose -f docker-compose.test.yml --env-file test.env exec -T postgres-test pg_isready -U postgres > /dev/null 2>&1; then \
			break; \
		fi; \
		echo "   Waiting... ($$timeout)"; \
		sleep 1; \
		timeout=$$((timeout - 1)); \
	done; \
	if [ $$timeout -eq 0 ]; then \
		echo "❌ Timeout waiting for postgres-test"; \
		docker-compose -f docker-compose.test.yml --env-file test.env down; \
		exit 1; \
	fi
	@echo "✅ postgres-test is ready"
	@echo "🧪 Running Go unit tests..."
	@docker-compose -f docker-compose.test.yml --env-file test.env run --rm \
		-e CGO_ENABLED=1 \
		back-test \
		sh -c "cd /app && go test -v -race -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html && echo '✅ Coverage report generated: coverage.html'" || \
		(echo "❌ Tests failed" && docker-compose -f docker-compose.test.yml --env-file test.env down && exit 1)
	@echo "🧹 Cleaning up test services..."
	@docker-compose -f docker-compose.test.yml --env-file test.env down
	@echo "✅ Unit tests completed!"

# Тесты бекенда (локально, без Docker)
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
	set -o allexport && source .env && set +o allexport && \
	docker-compose run --rm migrate \
		-path /migrations \
		-database "postgres://$$DB_USER:$$DB_PASSWORD@postgres:$$DB_PORT/$$DB_NAME?sslmode=disable" up

migrate-up-test:
	set -o allexport && source test.env && set +o allexport && \
	docker-compose run --rm migrate \
		-path /migrations \
		-database "postgres://$$DB_USER:$$DB_PASSWORD@postgres:$$DB_PORT/$$DB_NAME?sslmode=disable" up


migrate-down:
	docker-compose run --rm migrate -path /migrations -database "postgres://${DB_USER:-postgres}:${DB_PASSWORD:-password}@postgres:${DB_PORT:-5432}/${DB_NAME:-sumb}?sslmode=disable" 	down

# Помощь
help:
	@echo "Available commands:"
	@echo ""
	@echo "Development:"
	@echo "  make dev           - Start full development environment"
	@echo "  make dev-back      - Start dependencies (DB) in Docker, run Go app natively (for debugger)"
	@echo "  make back-debug    - Run Go app in Docker with Delve (remote debug)"
	@echo "  make dev-back-down - Stop dependencies for dev-back"
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
	@echo "Testing:"
	@echo "  make integration-back-tests - Run backend integration tests in Docker container"
	@echo "  make unit-back-tests - Run backend unit tests in Docker container"
	@echo "  make test-back     - Run Go tests locally (without Docker)"
	@echo "  make test-front    - Run Vue tests"
	@echo "  make init-env-test - Initialize test.env file from test.env.dist"
	@echo ""
	@echo "Utilities:"
	@echo "  make down          - Stop docker containers"
	@echo "  make logs          - View docker logs"

# По умолчанию показываем помощь
.DEFAULT_GOAL := help
