package adapter

import (
	"testing"

	generated "github.com/alexinator1/sumb/back/internal/domain/business/api/v1/businessgenerated"
	employee "github.com/alexinator1/sumb/back/internal/domain/employee/entity"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestRequestToDomainBusiness(t *testing.T) {
	t.Run("maps all fields correctly", func(t *testing.T) {
		description := "Test business description"
		middleName := "Ivanovich"

		req := generated.CreateBusinessRequest{
			Name:            "Test Business",
			Description:     &description,
			OwnerFirstName:  "Ivan",
			OwnerLastName:   "Petrov",
			OwnerMiddleName: &middleName,
			OwnerEmail:      openapi_types.Email("ivan.petrov@example.com"),
			OwnerPhone:      "+79990000000",
			Password:        "password123",
			PasswordConfirmation: "password123",
		}

		got, err := RequestToDomainBusiness(req)

		require.NoError(t, err)
		require.NotNil(t, got)
		require.Equal(t, "Test Business", got.Name)
		require.NotNil(t, got.Description)
		require.Equal(t, description, *got.Description)
		require.Equal(t, "Ivan", got.OwnerFirstName)
		require.Equal(t, "Petrov", got.OwnerLastName)
		require.NotNil(t, got.OwnerMiddleName)
		require.Equal(t, middleName, *got.OwnerMiddleName)
		require.Equal(t, "ivan.petrov@example.com", got.OwnerEmail)
		require.Equal(t, "+79990000000", got.OwnerPhone)
	})

	t.Run("maps fields with nil optional values", func(t *testing.T) {
		req := generated.CreateBusinessRequest{
			Name:            "Test Business",
			Description:     nil,
			OwnerFirstName:  "Ivan",
			OwnerLastName:   "Petrov",
			OwnerMiddleName: nil,
			OwnerEmail:      openapi_types.Email("ivan.petrov@example.com"),
			OwnerPhone:      "+79990000000",
			Password:        "password123",
			PasswordConfirmation: "password123",
		}

		got, err := RequestToDomainBusiness(req)

		require.NoError(t, err)
		require.NotNil(t, got)
		require.Equal(t, "Test Business", got.Name)
		require.Nil(t, got.Description)
		require.Equal(t, "Ivan", got.OwnerFirstName)
		require.Equal(t, "Petrov", got.OwnerLastName)
		require.Nil(t, got.OwnerMiddleName)
		require.Equal(t, "ivan.petrov@example.com", got.OwnerEmail)
		require.Equal(t, "+79990000000", got.OwnerPhone)
	})

	t.Run("converts email type correctly", func(t *testing.T) {
		req := generated.CreateBusinessRequest{
			Name:            "Test Business",
			OwnerFirstName:  "Ivan",
			OwnerLastName:   "Petrov",
			OwnerEmail:      openapi_types.Email("test@example.com"),
			OwnerPhone:      "+79990000000",
			Password:        "password123",
			PasswordConfirmation: "password123",
		}

		got, err := RequestToDomainBusiness(req)

		require.NoError(t, err)
		require.NotNil(t, got)
		require.Equal(t, "Test Business", got.Name)
		// Описание, middleName опциональны - могут быть nil
		require.Nil(t, got.Description)
		require.Equal(t, "Ivan", got.OwnerFirstName)
		require.Equal(t, "Petrov", got.OwnerLastName)
		require.Nil(t, got.OwnerMiddleName)
		require.Equal(t, "test@example.com", got.OwnerEmail)
		require.Equal(t, "+79990000000", got.OwnerPhone)
	})
}

func TestRequestToDomainEmployee(t *testing.T) {
	t.Run("maps all fields correctly and hashes password", func(t *testing.T) {
		middleName := "Ivanovich"
		password := "securePassword123"

		req := generated.CreateBusinessRequest{
			Name:            "Test Business",
			OwnerFirstName:  "Ivan",
			OwnerLastName:   "Petrov",
			OwnerMiddleName: &middleName,
			OwnerEmail:      openapi_types.Email("ivan.petrov@example.com"),
			OwnerPhone:      "+79990000000",
			Password:        password,
			PasswordConfirmation: password,
		}

		got, err := RequestToDomainEmployee(req)

		require.NoError(t, err)
		require.NotNil(t, got)
		require.Equal(t, "Ivan", got.FirstName)
		require.Equal(t, "Petrov", got.LastName)
		require.NotNil(t, got.MiddleName)
		require.Equal(t, middleName, *got.MiddleName)
		require.Equal(t, "+79990000000", got.Phone)
		require.Equal(t, "ivan.petrov@example.com", got.Email)
		require.Equal(t, employee.Owner, got.Role)
		require.NotEmpty(t, got.PasswordHash)

		// Verify password hash can be checked
		err = bcrypt.CompareHashAndPassword([]byte(got.PasswordHash), []byte(password))
		require.NoError(t, err)
	})

	t.Run("maps fields with nil optional values", func(t *testing.T) {
		password := "securePassword123"

		req := generated.CreateBusinessRequest{
			Name:            "Test Business",
			OwnerFirstName:  "Ivan",
			OwnerLastName:   "Petrov",
			OwnerMiddleName: nil,
			OwnerEmail:      openapi_types.Email("ivan.petrov@example.com"),
			OwnerPhone:      "+79990000000",
			Password:        password,
			PasswordConfirmation: password,
		}

		got, err := RequestToDomainEmployee(req)

		require.NoError(t, err)
		require.NotNil(t, got)
		require.Equal(t, "Ivan", got.FirstName)
		require.Equal(t, "Petrov", got.LastName)
		require.Nil(t, got.MiddleName)
		require.Equal(t, "+79990000000", got.Phone)
		require.Equal(t, "ivan.petrov@example.com", got.Email)
		require.Equal(t, employee.Owner, got.Role)
		require.NotEmpty(t, got.PasswordHash)

		// Verify password hash can be checked
		err = bcrypt.CompareHashAndPassword([]byte(got.PasswordHash), []byte(password))
		require.NoError(t, err)
	})

	t.Run("converts email type correctly", func(t *testing.T) {
		password := "securePassword123"

		req := generated.CreateBusinessRequest{
			Name:            "Test Business",
			OwnerFirstName:  "Ivan",
			OwnerLastName:   "Petrov",
			OwnerEmail:      openapi_types.Email("test@example.com"),
			OwnerPhone:      "+79990000000",
			Password:        password,
			PasswordConfirmation: password,
		}

		got, err := RequestToDomainEmployee(req)

		require.NoError(t, err)
		require.NotNil(t, got)
		require.Equal(t, "test@example.com", got.Email)
	})

	t.Run("generates different hashes for same password", func(t *testing.T) {
		password := "securePassword123"

		req := generated.CreateBusinessRequest{
			Name:            "Test Business",
			OwnerFirstName:  "Ivan",
			OwnerLastName:   "Petrov",
			OwnerEmail:      openapi_types.Email("ivan.petrov@example.com"),
			OwnerPhone:      "+79990000000",
			Password:        password,
			PasswordConfirmation: password,
		}

		got1, err1 := RequestToDomainEmployee(req)
		require.NoError(t, err1)

		got2, err2 := RequestToDomainEmployee(req)
		require.NoError(t, err2)

		// Hashes should be different (bcrypt uses salt)
		require.NotEqual(t, got1.PasswordHash, got2.PasswordHash)

		// But both should verify against the original password
		err1 = bcrypt.CompareHashAndPassword([]byte(got1.PasswordHash), []byte(password))
		require.NoError(t, err1)

		err2 = bcrypt.CompareHashAndPassword([]byte(got2.PasswordHash), []byte(password))
		require.NoError(t, err2)
	})

	t.Run("always sets role to Owner", func(t *testing.T) {
		password := "securePassword123"

		req := generated.CreateBusinessRequest{
			Name:            "Test Business",
			OwnerFirstName:  "Ivan",
			OwnerLastName:   "Petrov",
			OwnerEmail:      openapi_types.Email("ivan.petrov@example.com"),
			OwnerPhone:      "+79990000000",
			Password:        password,
			PasswordConfirmation: password,
		}

		got, err := RequestToDomainEmployee(req)

		require.NoError(t, err)
		require.NotNil(t, got)
		require.Equal(t, employee.Owner, got.Role)
		require.Equal(t, "Ivan", got.FirstName)
		require.Equal(t, "Petrov", got.LastName)
		require.Nil(t, got.MiddleName)
		require.Equal(t, "+79990000000", got.Phone)
		require.Equal(t, "ivan.petrov@example.com", got.Email)
		require.NotEmpty(t, got.PasswordHash)

		// Verify password hash can be checked
		err = bcrypt.CompareHashAndPassword([]byte(got.PasswordHash), []byte(password))
		require.NoError(t, err)	
	})
}

