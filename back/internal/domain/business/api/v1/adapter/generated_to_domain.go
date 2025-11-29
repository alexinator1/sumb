package adapter

import (
	"fmt"
	generated "github.com/alexinator1/sumb/back/internal/domain/business/api/v1/businessgenerated"
	domain "github.com/alexinator1/sumb/back/internal/domain/business/entity"
	employee "github.com/alexinator1/sumb/back/internal/domain/employee/entity"
	"golang.org/x/crypto/bcrypt"
	)

func RequestToDomainBusiness(req generated.CreateBusinessRequest) (*domain.Business, error) {

	b := &domain.Business{
		Name:            req.Name,
		Description:     req.Description,
		OwnerFirstName:  req.OwnerFirstName,
		OwnerLastName:   req.OwnerLastName,
		OwnerMiddleName: req.OwnerMiddleName,
		OwnerEmail:      string(req.OwnerEmail),
		OwnerPhone:      req.OwnerPhone,
	}
	return b, nil
}

func RequestToDomainEmployee(req generated.CreateBusinessRequest) (*employee.Employee, error) {
	
	
	bytesPassword := []byte(req.Password)

	passwordHash, err := bcrypt.GenerateFromPassword(bytesPassword, bcrypt.DefaultCost)
	if err != nil || passwordHash == nil {
		return nil, fmt.Errorf("failed to generate password hash: %w", err)
	}
	
	return &employee.Employee{
		FirstName:  req.OwnerFirstName,
		LastName:   req.OwnerLastName,
		MiddleName: req.OwnerMiddleName,
		Phone:      req.OwnerPhone,
		Email:      string(req.OwnerEmail),
		Role:       employee.Owner,
		PasswordHash: string(passwordHash),
	}, nil
}