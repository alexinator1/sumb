package employee

import "gorm.io/gorm"

type EmployeeProvider struct {
	db *gorm.DB
}

func NewEmployeeProvider(db *gorm.DB) *EmployeeProvider {
	return &EmployeeProvider{db: db}
}

func (p *EmployeeProvider) getEmployeeRoutes() {

}
