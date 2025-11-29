package adapter

import (
	"testing"
	"time"

	generated "github.com/alexinator1/sumb/back/internal/domain/business/api/v1/businessgenerated"
	domain "github.com/alexinator1/sumb/back/internal/domain/business/entity"
	"github.com/stretchr/testify/require"
)

func TestDomainToGeneratedBusiness(t *testing.T) {
	t.Run("returns nil when business is nil", func(t *testing.T) {
		require.Nil(t, DomainToGeneratedBusiness(nil))
	})

	t.Run("maps all available fields", func(t *testing.T) {
		now := time.Now()
		createdAt := now.Add(-96 * time.Hour)
		updatedAt := now
		deletedAt := now.Add(-24 * time.Hour)
		description := "Test business description"
		middleName := "Ivanovich"
		logoID := "logo-123"
		ownerID := uint64(42)

		input := &domain.Business{
			ID:              1,
			Name:            "Test Business",
			Description:     &description,
			OwnerFirstName:  "Ivan",
			OwnerLastName:   "Petrov",
			OwnerMiddleName: &middleName,
			OwnerEmail:      "ivan.petrov@example.com",
			OwnerPhone:      "+79990000000",
			LogoID:          &logoID,
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
			IsWorking:       true,
			DeletedAt:       &deletedAt,
			OwnerID:         &ownerID,
		}

		got := DomainToGeneratedBusiness(input)

		require.NotNil(t, got)
		require.IsType(t, (*generated.Business)(nil), got)
		require.Equal(t, uint64(1), got.Id)
		require.Equal(t, "Test Business", got.Name)
		require.NotNil(t, got.Description)
		require.Equal(t, description, *got.Description)
		require.Equal(t, "Ivan", got.OwnerFirstName)
		require.Equal(t, "Petrov", got.OwnerLastName)
		require.NotNil(t, got.OwnerMiddleName)
		require.Equal(t, middleName, *got.OwnerMiddleName)
		require.Equal(t, "ivan.petrov@example.com", string(got.OwnerEmail))
		require.Equal(t, "+79990000000", got.OwnerPhone)
		require.NotNil(t, got.LogoId)
		require.Equal(t, logoID, *got.LogoId)
		require.Equal(t, createdAt, got.CreatedAt)
		require.Equal(t, updatedAt, got.UpdatedAt)
		require.NotNil(t, got.IsWorking)
		require.Equal(t, true, *got.IsWorking)
		require.NotNil(t, got.DeletedAt)
		require.Equal(t, deletedAt, *got.DeletedAt)
		require.NotNil(t, got.OwnerId)
		require.Equal(t, ownerID, *got.OwnerId)
	})

	t.Run("maps fields with nil optional values", func(t *testing.T) {
		now := time.Now()
		createdAt := now.Add(-96 * time.Hour)
		updatedAt := now

		input := &domain.Business{
			ID:              2,
			Name:            "Test Business 2",
			Description:     nil,
			OwnerFirstName:  "Petr",
			OwnerLastName:   "Ivanov",
			OwnerMiddleName: nil,
			OwnerEmail:      "petr.ivanov@example.com",
			OwnerPhone:      "+79991111111",
			LogoID:          nil,
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
			IsWorking:       false,
			DeletedAt:       nil,
			OwnerID:         nil,
		}

		got := DomainToGeneratedBusiness(input)

		require.NotNil(t, got)
		require.Equal(t, uint64(2), got.Id)
		require.Equal(t, "Test Business 2", got.Name)
		require.Nil(t, got.Description)
		require.Equal(t, "Petr", got.OwnerFirstName)
		require.Equal(t, "Ivanov", got.OwnerLastName)
		require.Nil(t, got.OwnerMiddleName)
		require.Equal(t, "petr.ivanov@example.com", string(got.OwnerEmail))
		require.Equal(t, "+79991111111", got.OwnerPhone)
		require.Nil(t, got.LogoId)
		require.Equal(t, createdAt, got.CreatedAt)
		require.Equal(t, updatedAt, got.UpdatedAt)
		require.NotNil(t, got.IsWorking)
		require.Equal(t, false, *got.IsWorking)
		require.Nil(t, got.DeletedAt)
		require.Nil(t, got.OwnerId)
	})

	t.Run("converts email type correctly", func(t *testing.T) {
		now := time.Now()
		createdAt := now.Add(-96 * time.Hour)
		updatedAt := now

		input := &domain.Business{
			ID:              3,
			Name:            "Test Business 3",
			OwnerFirstName:  "Test",
			OwnerLastName:   "User",
			OwnerEmail:      "test@example.com",
			OwnerPhone:      "+79992222222",
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
			IsWorking:       true,
		}

		got := DomainToGeneratedBusiness(input)

		require.NotNil(t, got)
		require.Equal(t, "test@example.com", string(got.OwnerEmail))
	})

	t.Run("always converts IsWorking to pointer", func(t *testing.T) {
		now := time.Now()
		createdAt := now.Add(-96 * time.Hour)
		updatedAt := now

		testCases := []struct {
			name      string
			isWorking bool
		}{
			{"true", true},
			{"false", false},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				input := &domain.Business{
					ID:             4,
					Name:           "Test Business",
					OwnerFirstName: "Test",
					OwnerLastName:  "User",
					OwnerEmail:     "test@example.com",
					OwnerPhone:     "+79990000000",
					CreatedAt:      createdAt,
					UpdatedAt:      updatedAt,
					IsWorking:      tc.isWorking,
				}

				got := DomainToGeneratedBusiness(input)

				require.NotNil(t, got)
				require.NotNil(t, got.IsWorking)
				require.Equal(t, tc.isWorking, *got.IsWorking)
			})
		}
	})

	t.Run("converts ID to uint64", func(t *testing.T) {
		now := time.Now()
		createdAt := now.Add(-96 * time.Hour)
		updatedAt := now

		input := &domain.Business{
			ID:              999999,
			Name:            "Test Business",
			OwnerFirstName:  "Test",
			OwnerLastName:   "User",
			OwnerEmail:      "test@example.com",
			OwnerPhone:      "+79990000000",
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
			IsWorking:       true,
		}

		got := DomainToGeneratedBusiness(input)

		require.NotNil(t, got)
		require.Equal(t, uint64(999999), got.Id)
	})
}

