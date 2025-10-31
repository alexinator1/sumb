package app

import (
	"fmt"

	"github.com/alexinator1/sumb/back/internal/domain/employee/provider"
)

type AppProviders struct {
	cfg                 *AppConfig
	employeeApiProvider *provider.EmployeeApiProvider
}

func NewAppProviders(cfg *AppConfig) *AppProviders {
	return &AppProviders{cfg: cfg}
}

func (ap *AppProviders) Init() error {
	dbProvider, err := newDBProvider(ap.cfg)
	if err != nil {
		return fmt.Errorf("failed to initialize db provider: %w", err)
	}

	ap.employeeApiProvider = provider.NewEmployeeApiProvider(dbProvider.DB())

	return nil
}

func (ap *AppProviders) EmployeeApiProvider() *provider.EmployeeApiProvider {
	return ap.employeeApiProvider
}
