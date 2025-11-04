package adapter

import (
	generated "github.com/alexinator1/sumb/back/internal/domain/business/api/v1/businessgenerated"
	domain "github.com/alexinator1/sumb/back/internal/domain/business/entity"
	convertor "github.com/alexinator1/sumb/back/internal/tools/convertor"
)

func DomainToGeneratedBusiness(b *domain.Business) *generated.Business {
	if b == nil {
		return nil
	}
	return &generated.Business{
		Id:              int64(b.ID),
		Name:            b.Name,
		Description:     b.Description,
		OwnerFirstName:  b.OwnerFirstName,
		OwnerLastName:   b.OwnerLastName,
		OwnerMiddleName: b.OwnerMiddleName,
		OwnerEmail:      b.OwnerEmail,
		OwnerPhone:      b.OwnerPhone,
		LogoId:          b.LogoID,
		CreatedAt:       b.CreatedAt,
		UpdatedAt:       b.UpdatedAt,
		IsWorking:       b.IsWorking,
		DeletedAt:       b.DeletedAt,
		OwnerId:         b.OwnerID,
	}
}

func ptrToUint64OrZero(p *uint64) uint64 {
	if p == nil {
		return 0
	}
	return *p
}
