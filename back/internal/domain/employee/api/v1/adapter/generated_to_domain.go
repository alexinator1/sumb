package adapter

import (
	"fmt"

	generatedmodel "github.com/alexinator1/sumb/back/internal/domain/employee/api/v1/employeegenerated"
	domain "github.com/alexinator1/sumb/back/internal/domain/employee/entity"
	convertor "github.com/alexinator1/sumb/back/internal/tools/convertor"
)

func GeneratedToDomainEmployee(req generatedmodel.CreateEmployeeRequest) (*domain.Employee, error) {

	birthDate, err := convertor.PtrStringToTime(req.BirthDate)
	if err != nil {
		return nil, err
	}

	role := domain.EmployeeRole(req.Role)
	if !role.IsValid() {
		return nil, fmt.Errorf("invalid role: %s", req.Role)
	}

	return &domain.Employee{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		MiddleName: req.MiddleName,
		Phone:      req.Phone,
		Position:   req.Position,
		Role:       string(role),
		BirthDate:  birthDate,
		HiredAt:    req.HiredAt,
		Email:      req.Email,
		AvatarURL:  *req.AvatarUrl,
	}, nil
}
