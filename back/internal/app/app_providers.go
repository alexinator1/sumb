package app

import (
	"fmt"

	businessprovider "github.com/alexinator1/sumb/back/internal/domain/business/provider"
	employeprovider "github.com/alexinator1/sumb/back/internal/domain/employee/provider"
)

type AppProviders struct {
	cfg                 *AppConfig
	employeeApiProvider *employeprovider.EmployeeApiProvider
	businessApiProvider *businessprovider.BusinessApiProvider
}

func NewAppProviders(cfg *AppConfig) *AppProviders {
	return &AppProviders{cfg: cfg}
}

func (ap *AppProviders) Init() error {
	dbProvider, err := newDBProvider(ap.cfg)
	if err != nil {
		return fmt.Errorf("failed to initialize db provider: %w", err)
	}

	ap.employeeApiProvider = employeprovider.NewEmployeeApiProvider(dbProvider.DB())
	ap.businessApiProvider = businessprovider.NewBusinessApiProvider(dbProvider.DB())

	return nil
}

func (ap *AppProviders) EmployeeApiProvider() *employeprovider.EmployeeApiProvider {
	return ap.employeeApiProvider
}

func (ap *AppProviders) BusinessApiProvider() *businessprovider.BusinessApiProvider {
	return ap.businessApiProvider
}
