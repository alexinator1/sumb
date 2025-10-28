# Инфраструктурные Makefile

Эта папка содержит специализированные Makefile для различных аспектов проекта.

## Структура

- `modules.mk` - Генерация модульных API из OpenAPI спецификаций

## Модульная генерация (modules.mk)

### Назначение
Автоматическая генерация Go кода для отдельных API модулей из OpenAPI спецификаций.

### Использование

**Показать доступные модули:**
```bash
make list-modules
```

**Сгенерировать конкретный модуль:**
```bash
make generate-module MODULE=employee
make generate-module MODULE=sales
```

**Сгенерировать все модули:**
```bash
make generate-modules
```

**Добавить новый модуль:**
```bash
make add-module MODULE=newmodule
```

**Проверить спецификации:**
```bash
make check-module-specs
```

**Очистить все модули:**
```bash
make clean-modules
```

### Добавление нового модуля

1. **Добавьте название в массив MODULES** в `infra/makefiles/modules.mk`:
   ```makefile
   MODULES := employee sales inventory orders
   ```

2. **Создайте файл спецификации:**
   ```
   specs/api/v1/modules/inventory-api.yaml
   ```

3. **Сгенерируйте код:**
   ```bash
   make generate-module MODULE=inventory
   ```

### Структура результата

```
back/internal/api/v1/generated/
├── employee/     # Пакет employee
├── sales/        # Пакет sales  
├── inventory/    # Пакет inventory
└── orders/       # Пакет orders
```

### Использование в коде

```go
import (
    "back/internal/api/v1/generated/employee"
    "back/internal/api/v1/generated/sales"
    "back/internal/api/v1/generated/inventory"
)

// Использование Employee API
employeeService := employee.NewEmployeesAPIService()

// Использование Sales API  
salesService := sales.NewSalesAPIService()
```



