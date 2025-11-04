package adapter

import (
	generatedmodel "github.com/alexinator1/sumb/back/internal/domain/employee/api/v1/employeegenerated"
	domain "github.com/alexinator1/sumb/back/internal/domain/employee/entity"
	convertor "github.com/alexinator1/sumb/back/internal/tools/convertor"
)

func DomainToGeneratedEmployee(e *domain.Employee) *generatedmodel.Employee {	
	if e == nil {
		return nil
	}

	var birthDate *string
	if e.BirthDate != nil {
		d := e.BirthDate.Format("2006-01-02")
		birthDate = &d
	}

	return &generatedmodel.Employee{
		Id:           int64(e.ID),
		FirstName:    e.FirstName,
		LastName:     e.LastName,
		MiddleName:   e.MiddleName,
		Phone:        e.Phone,
		Position:     e.Position,
		Role:         e.Role,
		BirthDate:    birthDate,
		HiredAt:      e.HiredAt,
		FiredAt:      e.FiredAt,
		AddedAt:      e.AddedAt,
		Email:        e.Email,
		Status:       e.Status,
		CreatedAt:    e.CreatedAt,
		UpdatedAt:    e.UpdatedAt,
		CreatedBy:    convertor.Int64PtrIfNotZero(e.CreatedBy),
		AvatarUrl:    convertor.StrPtrIfNotEmpty(e.AvatarURL),
	}
}
