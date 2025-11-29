package helpers

import (
	"fmt"

	businessentity "github.com/alexinator1/sumb/back/internal/domain/business/entity"
	employeeentity "github.com/alexinator1/sumb/back/internal/domain/employee/entity"
	"github.com/oapi-codegen/runtime/types"
	"gorm.io/gorm"
)

type DbVerifier struct {
	db *gorm.DB
}

func NewDbVerifier(db *gorm.DB) *DbVerifier {
	return &DbVerifier{db: db}
}

func (v *DbVerifier) VerifyBusiness(businessName string) error {
	var count int64
	if err := v.db.Model(&businessentity.Business{}).Where("name = ?", businessName).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("business not found")
	}
	return nil
}

func (v *DbVerifier) VerifyEmployee(email types.Email) error {
	var count int64
	if err := v.db.Model(&employeeentity.Employee{}).Where("email = ?", string(email)).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("employee not found")
	}
	return nil
}
