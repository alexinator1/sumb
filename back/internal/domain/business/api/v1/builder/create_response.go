package builder

import (
	"github.com/alexinator1/sumb/back/internal/domain/business/entity"
	"github.com/alexinator1/sumb/back/internal/domain/business/api/v1/adapter"
	generated "github.com/alexinator1/sumb/back/internal/domain/business/api/v1/businessgenerated"
)

func BuildCreateResponse(business *entity.Business) generated.CreateBusinessResponse {
	return generated.CreateBusinessResponse{
		Data:    *adapter.DomainToGeneratedBusiness(business),
		Message: "Business created successfully",
	}
}