package employee

import (
	"context"
	"fmt"
)

// Service provides business logic operations on employees
type Service struct {
	repo *Repo
}

// NewService creates a new employee service
func NewService(repo *Repo) *Service {
	return &Service{
		repo: repo,
	}
}

// GetEmployeeByID retrieves an employee by their ID
func (s *Service) GetEmployeeByID(ctx context.Context, id uint64) (*Employee, error) {
	if id == 0 {
		return nil, fmt.Errorf("invalid employee ID: cannot be zero")
	}

	employee, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get employee by ID: %w", err)
	}

	return employee, nil
}
