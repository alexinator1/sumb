package adapter

import (
	generatedmodel "github.com/alexinator1/sumb/back/internal/domain/employee/api/v1/employeegenerated"
	domain "github.com/alexinator1/sumb/back/internal/domain/employee/entity"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func DomainToGeneratedEmployee(de *domain.Employee) *generatedmodel.Employee {
	if de == nil {
		return nil
	}

	if de.BirthDate == nil {
		return nil
	}
	role := generatedmodel.EmployeeRole(string(de.Role))
	status := generatedmodel.EmployeeStatus(string(de.Status))
	createBy := de.CreatedBy
	hiredAt := de.HiredAt
	firedAt := de.FiredAt
	addedAt := de.AddedAt
	createdAt := de.CreatedAt
	updatedAt := de.UpdatedAt
	avatarUrl := de.AvatarURL

	generated := generatedmodel.Employee{
		Id:         de.ID,
		FirstName:  de.FirstName,
		LastName:   de.LastName,
		MiddleName: de.MiddleName,
		Phone:      de.Phone,
		Position:   de.Position,
		Role:       &role,
		HiredAt:    hiredAt,
		FiredAt:    firedAt,
		AddedAt:    &addedAt,
		Email:      openapi_types.Email(de.Email),
		Status:     &status,
		CreatedAt:  &createdAt,
		UpdatedAt:  &updatedAt,
		CreatedBy:  createBy,
		AvatarUrl:  &avatarUrl,
	}
	if de.BirthDate != nil {
		birthDate := openapi_types.Date{
			Time: *de.BirthDate,
		}
		generated.BirthDate = &birthDate
	}
	return &generated
}
