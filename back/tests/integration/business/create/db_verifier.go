package create

import (
	"fmt"

	businessgenerated "github.com/alexinator1/sumb/back/internal/domain/business/api/v1/businessgenerated"
	businessentity "github.com/alexinator1/sumb/back/internal/domain/business/entity"
	employeeentity "github.com/alexinator1/sumb/back/internal/domain/employee/entity"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type DbVerifier struct {
	db      *gorm.DB
	require *require.Assertions
}

func NewDbVerifier(db *gorm.DB, require *require.Assertions) *DbVerifier {
	return &DbVerifier{db: db, require: require}
}

func (v *DbVerifier) VerifyBusiness(expectedCreateResponse businessgenerated.CreateBusinessResponse) error {
	var business businessentity.Business
	if err := v.db.Model(&businessentity.Business{}).Where("name = ?", expectedCreateResponse.Data.Name).First(&business).Error; err != nil {
		return fmt.Errorf("business не найден в бд")
	}
	v.require.Equal(expectedCreateResponse.Data.Id, business.ID, "business id does not match expected values")
	v.require.Equal(expectedCreateResponse.Data.Name, business.Name, "business name does not match expected values")
	v.require.Equal(*expectedCreateResponse.Data.IsWorking, business.IsWorking, "business is working does not match expected values")
	v.require.NotNil(business.CreatedAt, "business created at does not match expected values")
	v.require.NotNil(business.UpdatedAt, "business updated at does not match expected values")
	return nil
}

func (v *DbVerifier) VerifyOwner(expectedCreateResponse businessgenerated.CreateBusinessResponse) error {
	var employee employeeentity.Employee
	if err := v.db.Model(&employeeentity.Employee{}).Where("email = ?", string(expectedCreateResponse.Data.OwnerEmail)).First(&employee).Error; err != nil {
		return fmt.Errorf("failed to get employee by email: %w", err)
	}

	expectedIdOwnerId := *expectedCreateResponse.Data.OwnerId

	v.require.Equal(expectedCreateResponse.Data.OwnerFirstName, employee.FirstName, "employee first name does not match expected values")
	v.require.Equal(expectedCreateResponse.Data.OwnerLastName, employee.LastName, "employee last name does not match expected values")
	v.require.Equal(expectedCreateResponse.Data.OwnerPhone, employee.Phone, "employee phone does not match expected values")
	v.require.Equal(string(expectedCreateResponse.Data.OwnerEmail), employee.Email, "employee email does not match expected values")
	v.require.Equal(expectedIdOwnerId, employee.ID, "employee id does not match expected values")
	v.require.Equal(employeeentity.Owner, employee.Role, "employee role does not match expected values")
	v.require.NotNil(employee.PasswordHash, "employee password hash is nil")
	v.require.NotNil(employee.CreatedAt, "employee created at is nil")
	v.require.NotNil(employee.UpdatedAt, "employee updated at is nil")

	return nil
}

func (v *DbVerifier) VerifyBusinessAndOwner(expectedCreateResponse businessgenerated.CreateBusinessResponse) error {
	err := v.VerifyBusiness(expectedCreateResponse)
	if err != nil {
		return fmt.Errorf("failed to verify business: %w", err)
	}
	err = v.VerifyOwner(expectedCreateResponse)
	if err != nil {
		return fmt.Errorf("failed to verify employee: %w", err)
	}
	return nil
}

func (v *DbVerifier) VerifyEmptyTable(tableName string) error {
	var count int64
	if err := v.db.Table(tableName).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to count rows in table %s: %w", tableName, err)
	}
	v.require.Equal(int64(0), count)
	return nil
}
