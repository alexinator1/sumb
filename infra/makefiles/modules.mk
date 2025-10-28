# =============================================================================
# Модульная генерация API
# =============================================================================

# Массив модулей для генерации (добавляйте новые модули сюда)
MODULES := employee sales

# Переменные для генерации модулей
MODULES_OUTPUT := back/internal/api/v1/generated
MODULES_SPEC_DIR := specs/api/v1/modules
MODULES_GENERATED_FOLDER := 
MODULE_GEN_PATH = $(MODULES_OUTPUT)/$(MODULE)generated

# Создание полной спецификации для модуля
# Использование: make create-module-spec MODULE=employee
create-module-spec:
	@if [ -z "$(MODULE)" ]; then \
		echo "❌ Error: MODULE variable is required. Usage: make create-module-spec MODULE=module_name"; \
		echo "Available modules: $(MODULES)"; \
		exit 1; \
	fi
	@echo "🔧 Creating full specification for $(MODULE) module..."
	@mkdir -p $(MODULES_SPEC_DIR)/full
	@echo "openapi: 3.0.3" > $(MODULES_SPEC_DIR)/full/$(MODULE)-full.yaml
	@echo "info:" >> $(MODULES_SPEC_DIR)/full/$(MODULE)-full.yaml
	@echo "  title: $(MODULE) API" >> $(MODULES_SPEC_DIR)/full/$(MODULE)-full.yaml
	@echo "  version: 1.0.0" >> $(MODULES_SPEC_DIR)/full/$(MODULE)-full.yaml
	@echo "  description: API для модуля $(MODULE)" >> $(MODULES_SPEC_DIR)/full/$(MODULE)-full.yaml
	@echo "" >> $(MODULES_SPEC_DIR)/full/$(MODULE)-full.yaml
	@echo "servers:" >> $(MODULES_SPEC_DIR)/full/$(MODULE)-full.yaml
	@echo "  - url: http://localhost:8080/api/v1" >> $(MODULES_SPEC_DIR)/full/$(MODULE)-full.yaml
	@echo "    description: Development server" >> $(MODULES_SPEC_DIR)/full/$(MODULE)-full.yaml
	@echo "" >> $(MODULES_SPEC_DIR)/full/$(MODULE)-full.yaml
	@cat $(MODULES_SPEC_DIR)/$(MODULE)-api.yaml >> $(MODULES_SPEC_DIR)/full/$(MODULE)-full.yaml
	@echo "✅ Full specification created: $(MODULES_SPEC_DIR)/full/$(MODULE)-full.yaml"

# Универсальная функция для генерации модуля
# Использование: make generate-module MODULE=employee
generate-module:
	@if [ -z "$(MODULE)" ]; then \
		echo "❌ Error: MODULE variable is required. Usage: make generate-module MODULE=module_name"; \
		echo "Available modules: $(MODULES)"; \
		exit 1; \
	fi
	@echo "🔧 Creating full specification for $(MODULE)..."
	@$(MAKE) create-module-spec MODULE=$(MODULE)
	@echo "🔧 Generating $(MODULE) module code..."
	@echo "🧹 Cleaning previous generated code..."
	@rm -rf $(MODULE_GEN_PATH)
	@mkdir -p $(MODULE_GEN_PATH)
	@docker run --rm -v ${PWD}:/local \
		$(OPENAPI_GENERATOR) generate \
		-i /local/$(MODULES_SPEC_DIR)/full/$(MODULE)-full.yaml \
		-g go-server \
		-o /local/$(MODULE_GEN_PATH) \
		--additional-properties=packageName=$(MODULE)generated,enumClassPrefix=true,withGoCodegenComment=true \
		--skip-validate-spec
	@echo "✅ $(MODULE) module code generated in $(MODULES_OUTPUT)/$(MODULE)"
	@echo "📁 Moving Go files from go/ subdirectory..."
	@if [ -d "$(MODULE_GEN_PATH)/go" ]; then \
		mv $(MODULE_GEN_PATH)/go/*.go $(MODULE_GEN_PATH)/; \
		rmdir $(MODULE_GEN_PATH)/go; \
		echo "✅ Go files moved to $(MODULE_GEN_PATH)"; \
	fi
	@echo "🧹 Cleaning up generated go.mod file..."
	@if [ -f "$(MODULE_GEN_PATH)/go.mod" ]; then \
		rm $(MODULE_GEN_PATH)/go.mod; \
		echo "✅ Removed separate go.mod file"; \
	fi
	@echo "🧹 Removing non-Go files, main.go and unnecessary directories..."
	@find $(MODULE_GEN_PATH) -name "*.go" -exec basename {} \; | grep -v "main.go" > /tmp/go_files.txt; \
	find $(MODULE_GEN_PATH) -type f ! -name "*.go" -delete; \
	rm -f $(MODULE_GEN_PATH)/main.go; \
	rm -rf $(MODULE_GEN_PATH)/api; \
	rm -rf $(MODULE_GEN_PATH)/.openapi-generator; \
	echo "✅ Removed non-Go files, main.go and unnecessary directories"
	@echo "📁 Generated $(MODULE) Go files:"
	@find $(MODULE_GEN_PATH) -name "*.go" | head -10

# Генерация всех модулей из массива MODULES
generate-modules:
	@echo "🔧 Generating all modules: $(MODULES)"
	@for module in $(MODULES); do \
		echo "📦 Generating module: $$module"; \
		$(MAKE) generate-module MODULE=$$module; \
		echo ""; \
	done
	@echo "✅ All modules generated successfully!"

# Показать доступные модули
list-modules:
	@echo "📦 Available modules:"
	@for module in $(MODULES); do \
		echo "  - $$module ($(MODULES_SPEC_DIR)/$$module-api.yaml)"; \
	done
	@echo ""
	@echo "Usage: make generate-module MODULE=module_name"
	@echo "Example: make generate-module MODULE=employee"

# Добавить новый модуль в массив MODULES
# Использование: make add-module MODULE=newmodule
add-module:
	@if [ -z "$(MODULE)" ]; then \
		echo "❌ Error: MODULE variable is required. Usage: make add-module MODULE=module_name"; \
		exit 1; \
	fi
	@echo "📝 Adding module '$(MODULE)' to MODULES array..."
	@echo "Please add '$(MODULE)' to the MODULES array in infra/makefiles/modules.mk"
	@echo "Current modules: $(MODULES)"
	@echo "New modules should be: $(MODULES) $(MODULE)"
	@echo ""
	@echo "Also create the specification file: $(MODULES_SPEC_DIR)/$(MODULE)-api.yaml"
	@echo "Then run: make generate-module MODULE=$(MODULE)"

# Проверить существование файлов спецификаций для всех модулей
check-module-specs:
	@echo "🔍 Checking module specifications..."
	@for module in $(MODULES); do \
		if [ -f "$(MODULES_SPEC_DIR)/$$module-api.yaml" ]; then \
			echo "✅ $$module: $(MODULES_SPEC_DIR)/$$module-api.yaml"; \
		else \
			echo "❌ $$module: $(MODULES_SPEC_DIR)/$$module-api.yaml (missing)"; \
		fi; \
	done

# Очистить все сгенерированные модули
clean-modules:
	@echo "🧹 Cleaning all generated modules..."
	@for module in $(MODULES); do \
		if [ -d "$(MODULES_OUTPUT)/$$module" ]; then \
			rm -rf $(MODULES_OUTPUT)/$$module; \
			echo "✅ Removed $(MODULES_OUTPUT)/$$module"; \
		fi; \
	done
	@echo "✅ All modules cleaned"

# Очистить полные спецификации модулей
clean-module-specs:
	@echo "🧹 Cleaning full module specifications..."
	@if [ -d "$(MODULES_SPEC_DIR)/full" ]; then \
		rm -rf $(MODULES_SPEC_DIR)/full; \
		echo "✅ Removed $(MODULES_SPEC_DIR)/full"; \
	fi
	@echo "✅ Module specifications cleaned"

# Извлечь общие файлы в папку common
extract-common-files:
	@echo "📦 Extracting common files to common/ directory..."
	@mkdir -p $(MODULES_OUTPUT)/common
	@# Берем общие файлы из первого модуля (employee)
	@if [ -d "$(MODULES_OUTPUT)/employee" ]; then \
		if [ -f "$(MODULES_OUTPUT)/employee/error.go" ]; then \
			cp $(MODULES_OUTPUT)/employee/error.go $(MODULES_OUTPUT)/common/; \
			echo "✅ Copied error.go to common/"; \
		fi; \
		if [ -f "$(MODULES_OUTPUT)/employee/helpers.go" ]; then \
			cp $(MODULES_OUTPUT)/employee/helpers.go $(MODULES_OUTPUT)/common/; \
			echo "✅ Copied helpers.go to common/"; \
		fi; \
		if [ -f "$(MODULES_OUTPUT)/employee/impl.go" ]; then \
			cp $(MODULES_OUTPUT)/employee/impl.go $(MODULES_OUTPUT)/common/; \
			echo "✅ Copied impl.go to common/"; \
		fi; \
		if [ -f "$(MODULES_OUTPUT)/employee/logger.go" ]; then \
			cp $(MODULES_OUTPUT)/employee/logger.go $(MODULES_OUTPUT)/common/; \
			echo "✅ Copied logger.go to common/"; \
		fi; \
		if [ -f "$(MODULES_OUTPUT)/employee/routers.go" ]; then \
			cp $(MODULES_OUTPUT)/employee/routers.go $(MODULES_OUTPUT)/common/; \
			echo "✅ Copied routers.go to common/"; \
		fi; \
		if [ -f "$(MODULES_OUTPUT)/employee/api.go" ]; then \
			cp $(MODULES_OUTPUT)/employee/api.go $(MODULES_OUTPUT)/common/; \
			echo "✅ Copied api.go to common/"; \
		fi; \
	fi
	@echo "✅ Common files extracted to $(MODULES_OUTPUT)/common"

# Удалить общие файлы из модулей
remove-common-from-modules:
	@echo "🧹 Removing common files from modules..."
	@for module in $(MODULES); do \
		if [ -d "$(MODULES_OUTPUT)/$$module" ]; then \
			rm -f $(MODULES_OUTPUT)/$$module/error.go; \
			rm -f $(MODULES_OUTPUT)/$$module/helpers.go; \
			rm -f $(MODULES_OUTPUT)/$$module/impl.go; \
			rm -f $(MODULES_OUTPUT)/$$module/logger.go; \
			rm -f $(MODULES_OUTPUT)/$$module/routers.go; \
			rm -f $(MODULES_OUTPUT)/$$module/api.go; \
			echo "✅ Removed common files from $$module"; \
		fi; \
	done
	@echo "✅ Common files removed from all modules"
