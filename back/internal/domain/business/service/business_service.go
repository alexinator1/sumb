package service

import (
	"context"
	"fmt"

	entity "github.com/alexinator1/sumb/back/internal/domain/business/entity"
	repo "github.com/alexinator1/sumb/back/internal/domain/business/repository"
	employeeEntity "github.com/alexinator1/sumb/back/internal/domain/employee/entity"
	employeeRepo "github.com/alexinator1/sumb/back/internal/domain/employee/repository"
	"gorm.io/gorm"
)

type BusinessService struct {
	repo    *repo.BusinessRepo
	empRepo *employeeRepo.EmployeeRepo
	db      *gorm.DB
}

func NewService(repo *repo.BusinessRepo, empRepo *employeeRepo.EmployeeRepo, db *gorm.DB) *BusinessService {
	return &BusinessService{
		repo:    repo,
		empRepo: empRepo,
		db:      db,
	}
}

func (bs *BusinessService) GetBusinessByID(ctx context.Context, id uint64) (*entity.Business, error) {
	if id == 0 {
		return nil, fmt.Errorf("invalid business ID: cannot be zero")
	}
	b, err := bs.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get business by ID: %w", err)
	}
	return b, nil
}

func (bs *BusinessService) CreateBusinessWithOwner(ctx context.Context, b *entity.Business, e *employeeEntity.Employee) (*entity.Business, error) {
	if b == nil {
		return nil, fmt.Errorf("business cannot be nil")
	}
	if b.Name == "" {
		return nil, fmt.Errorf("business name is required")
	}

	// Начинаем транзакцию
	tx := bs.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	// Откатываем транзакцию в случае ошибки
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	// Создаем бизнес в транзакции
	if err := bs.repo.CreateWithDB(ctx, tx, b); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create business: %w", err)
	}

	// Создаем сотрудника в транзакции
	e.BusinessID = b.ID
	if err := bs.empRepo.CreateWithDB(ctx, tx, e); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create owner: %w", err)
	}

	// Обновляем бизнес с ID владельца в транзакции
	b.OwnerID = &e.ID
	if err := bs.repo.UpdateWithDB(ctx, tx, b); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to attach owner to business: %w", err)
	}

	// Фиксируем транзакцию
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return b, nil
}
