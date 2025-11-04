package business

import (
	"context"
	"fmt"

	entity "github.com/alexinator1/sumb/back/internal/domain/business/entity"
	repo "github.com/alexinator1/sumb/back/internal/domain/business/repository"
	employeeEntity "github.com/alexinator1/sumb/back/internal/domain/employee/entity"
	employeeService "github.com/alexinator1/sumb/back/internal/domain/employee/service"
)

type BusinessService struct {
	repo *repo.BusinessRepo
	es   *employeeService.EmployeeService
}

func NewService(repo *repo.BusinessRepo, es *employeeService.EmployeeService) *BusinessService {
	return &BusinessService{
		repo: repo,
		es:   es,
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

	// Set default values
	if !b.IsWorking {
		b.IsWorking = true
	}

	if err := bs.repo.Create(ctx, b); err != nil {
		return nil, fmt.Errorf("failed to create business: %w", err)
	}

	e.BusinessID = b.ID
	if err := bs.es.Create(ctx, e); err != nil {
		return nil, fmt.Errorf("failed to create owner: %w", err)
	}

	b.OwnerID = &e.ID
	if err := bs.repo.Update(ctx, b); err != nil {
		return nil, fmt.Errorf("failed to attach owner to business: %w", err)
	}
	
	return b, nil	
}
