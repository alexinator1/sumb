package app

import (
	"fmt"

	businessprovider "github.com/alexinator1/sumb/back/internal/domain/business/provider"
	employeprovider "github.com/alexinator1/sumb/back/internal/domain/employee/provider"
)

type AppProviders struct {
	cfg                 *AppConfig
	DbProvider          *DBProvider
	employeeProvider *employeprovider.EmployeeProvider
	businessProvider *businessprovider.BusinessProvider
}

func NewAppProviders(cfg *AppConfig) (*AppProviders, error) {
	dbProvider, err := newDBProvider(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize db provider: %w", err)
	}
	return &AppProviders{cfg: cfg, DbProvider: dbProvider}, nil
}

func (ap *AppProviders) Init() error {
	ap.employeeProvider = employeprovider.NewEmployeeProvider(ap.DbProvider.DB())
	ap.businessProvider = businessprovider.NewBusinessProvider(ap.DbProvider.DB())
	return nil
}

func (ap *AppProviders) EmployeeProvider() *employeprovider.EmployeeProvider {
	if ap.employeeProvider == nil {
		ap.employeeProvider = employeprovider.NewEmployeeProvider(ap.DbProvider.DB())
	}

	return ap.employeeProvider
}

func (ap *AppProviders) BusinessProvider() *businessprovider.BusinessProvider {
	if ap.businessProvider == nil {
		ap.businessProvider = businessprovider.NewBusinessProvider(ap.DbProvider.DB())
	}

	return ap.businessProvider
}