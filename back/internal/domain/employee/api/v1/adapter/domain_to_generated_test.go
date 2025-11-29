package adapter

import (
	"testing"
	"time"

	generatedmodel "github.com/alexinator1/sumb/back/internal/domain/employee/api/v1/employeegenerated"
	domain "github.com/alexinator1/sumb/back/internal/domain/employee/entity"
	"github.com/stretchr/testify/require"
)

func TestDomainToGeneratedEmployee(t *testing.T) {
	t.Run("returns nil when employee is nil", func(t *testing.T) {
		require.Nil(t, DomainToGeneratedEmployee(nil))
	})

	t.Run("returns nil when birth date is missing", func(t *testing.T) {
		require.Nil(t, DomainToGeneratedEmployee(&domain.Employee{}))
	})

	t.Run("maps all available fields", func(t *testing.T) {
		now := time.Now()
		birthDate := now.AddDate(-25, 0, 0)
		hiredAt := now.Add(-48 * time.Hour)
		firedAt := now.Add(48 * time.Hour)
		addedAt := now.Add(-72 * time.Hour)
		createdAt := now.Add(-96 * time.Hour)
		updatedAt := now
		createdBy := uint64(99)
		avatarURL := "https://cdn.example.com/avatar.png"

		input := &domain.Employee{
			ID:         1,
			FirstName:  "Ivan",
			LastName:   "Petrov",
			MiddleName: strPtr("Ivanovich"),
			Phone:      "+79990000000",
			Position:   strPtr("Founder"),
			Role:       domain.Owner,
			BirthDate:  &birthDate,
			HiredAt:    &hiredAt,
			FiredAt:    &firedAt,
			AddedAt:    addedAt,
			Email:      "ivan.petrov@example.com",
			Status:     "active",
			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
			CreatedBy:  &createdBy,
			AvatarURL:  avatarURL,
		}

		got := DomainToGeneratedEmployee(input)

		require.NotNil(t, got)
		require.Equal(t, uint64(1), got.Id)
		require.Equal(t, "Ivan", got.FirstName)
		require.Equal(t, "Petrov", got.LastName)
		require.NotNil(t, got.MiddleName)
		require.Equal(t, "Ivanovich", *got.MiddleName)
		require.Equal(t, "+79990000000", got.Phone)

		require.NotNil(t, got.Position)
		require.Equal(t, "Founder", *got.Position)

		require.NotNil(t, got.Role)
		require.Equal(t, generatedmodel.EmployeeRole(input.Role), *got.Role)

		require.NotNil(t, got.Status)
		require.Equal(t, generatedmodel.EmployeeStatus(input.Status), *got.Status)

		require.NotNil(t, got.BirthDate)
		require.Equal(t, birthDate, got.BirthDate.Time)

		require.NotNil(t, got.HiredAt)
		require.Equal(t, hiredAt, *got.HiredAt)

		require.NotNil(t, got.FiredAt)
		require.Equal(t, firedAt, *got.FiredAt)

		require.NotNil(t, got.AddedAt)
		require.Equal(t, addedAt, *got.AddedAt)

		require.NotNil(t, got.CreatedAt)
		require.Equal(t, createdAt, *got.CreatedAt)

		require.NotNil(t, got.UpdatedAt)
		require.Equal(t, updatedAt, *got.UpdatedAt)

		require.NotNil(t, got.CreatedBy)
		require.Equal(t, createdBy, *got.CreatedBy)

		require.NotNil(t, got.AvatarUrl)
		require.Equal(t, avatarURL, *got.AvatarUrl)

		require.Equal(t, input.Email, string(got.Email))
	})
}

func strPtr[T any](value T) *T {
	return &value
}
