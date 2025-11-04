package adapter

import (
	generated "github.com/alexinator1/sumb/back/internal/domain/business/api/v1/businessgenerated"
	domain "github.com/alexinator1/sumb/back/internal/domain/business/entity"
	employee "github.com/alexinator1/sumb/back/internal/domain/employee/entity"
)

func RequestToDomainBusiness(req generated.CreateBusinessRequest) (*domain.Business, error) {
	b := &domain.Business{
		Name:            req.Name,
		Description:     req.Description,
		OwnerFirstName:  req.OwnerFirstName,
		OwnerLastName:   req.OwnerLastName,
		OwnerMiddleName: req.OwnerMiddleName,
		OwnerEmail:      req.OwnerEmail,
		OwnerPhone:      req.OwnerPhone,
		LogoID:          req.LogoId,
		IsWorking:       req.IsWorking,
	}
	return b, nil
}

func RequestToDomainEmployee(req generated.CreateBusinessRequest) (*employee.Employee, error) {
	e := &employee.Employee{
		FirstName:  req.OwnerFirstName,
		LastName:   req.OwnerLastName,
		MiddleName: req.OwnerMiddleName,
		Phone:      req.OwnerPhone,
		Email:      req.OwnerEmail,
		Role:       employee.Owner,
		BusinessID: uint64(*req.OwnerId),
	}
	return e, nil
}
