package employee

import (
	"context"
	"fmt"
	entity "github.com/alexinator1/sumb/back/internal/domain/employee/entity"
	repo "github.com/alexinator1/sumb/back/internal/domain/employee/repository"
)

// EmployeeService provides business logic operations on employees
type EmployeeService struct {
	repo *repo.EmployeeRepo
}

// NewService creates a new employee service
func NewService(repo *repo.EmployeeRepo) *EmployeeService {
	return &EmployeeService{
		repo: repo,
	}
}

// GetEmployeeByID retrieves an employee by their ID
func (s *EmployeeService) GetEmployeeByID(ctx context.Context, id uint64) (*entity.Employee, error) {
	if id == 0 {
		return nil, fmt.Errorf("invalid employee ID: cannot be zero")
	}

	employee, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get employee by ID: %w", err)
	}

	return employee, nil
}

func (s *EmployeeService) Create(ctx context.Context, e *entity.Employee) error {
	if e == nil {
		return fmt.Errorf("employee cannot be nil")
	}
	if e.FirstName == "" {
		return fmt.Errorf("employee first name is required")
	}
	if e.LastName == "" {
		return fmt.Errorf("employee last name is required")
	}
	if err := s.repo.Create(ctx, e); err != nil {
		return fmt.Errorf("failed to create employee: %w", err)
	}
	return nil
}
