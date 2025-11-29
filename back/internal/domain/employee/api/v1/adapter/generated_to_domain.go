package adapter

import (
	"fmt"

	generatedmodel "github.com/alexinator1/sumb/back/internal/domain/employee/api/v1/employeegenerated"
	domain "github.com/alexinator1/sumb/back/internal/domain/employee/entity"
)

func GeneratedToDomainEmployee(req generatedmodel.CreateEmployeeRequest) (*domain.Employee, error) {
	if req.BirthDate == nil {
		return nil, fmt.Errorf("birth date is required")
	}
	birthDate := req.BirthDate.Time

	if req.Role == nil {
		return nil, fmt.Errorf("role is required")
	}
	role := domain.EmployeeRole(*req.Role)

	var avatarURL string
	if req.AvatarUrl != nil {
		avatarURL = *req.AvatarUrl
	}

	return &domain.Employee{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		MiddleName: req.MiddleName,
		Phone:      req.Phone,
		Position:   req.Position,
		Role:       role,
		BirthDate:  &birthDate,
		HiredAt:    req.HiredAt,
		Email:      string(req.Email),
		AvatarURL:  avatarURL,
	}, nil
}
