package app

import (
	"fmt"

	"github.com/alexinator1/sumb/back/internal/employee"
)

type AppProviders struct {
	cfg              *AppConfig
	employeeProvider *employee.EmployeeProvider
}

func NewAppProviders(cfg *AppConfig) *AppProviders {
	return &AppProviders{cfg: cfg}
}

func (ap *AppProviders) Init() error {
	dbProvider, err := NewDBProvider(ap.cfg)
	if err != nil {
		return fmt.Errorf("failed to initialize db provider: %w", err)
	}

	ap.employeeProvider = employee.NewEmployeeProvider(dbProvider.DB())

	return nil
}

func (ap *AppProviders) EmployeeProvider() *employee.EmployeeProvider {
	return ap.employeeProvider
}
