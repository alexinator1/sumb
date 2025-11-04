package adapter

import (
	"testing"
	"time"

	domain "github.com/alexinator1/sumb/back/internal/domain/employee/entity"
)

func TestDomainToGeneratedEmployee_NilInput(t *testing.T) {
	result := DomainToGeneratedEmployee(nil)
	if result != nil {
		t.Errorf("Expected nil for nil input, got %+v", result)
	}
}

func TestDomainToGeneratedEmployee_FullEmployee(t *testing.T) {
	birthDate := time.Date(1990, 5, 20, 0, 0, 0, 0, time.UTC)
	hiredAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	firedAt := time.Date(2025, 6, 1, 18, 0, 0, 0, time.UTC)
	addedAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	updatedAt := time.Date(2024, 1, 16, 12, 0, 0, 0, time.UTC)

	middleName := "Иванович"
	position := "Кассир"
	avatarURL := "https://cdn.example.com/avatars/1.png"

	domainEmp := &domain.Employee{
		ID:           123,
		FirstName:    "Иван",
		LastName:     "Петров",
		MiddleName:   &middleName,
		PasswordHash: "$2b$12$...",
		Phone:        "+79161234567",
		Position:     &position,
		Role:         "regular",
		BirthDate:    &birthDate,
		HiredAt:      &hiredAt,
		FiredAt:      &firedAt,
		AddedAt:      addedAt,
		Email:        "ivan.petrov@example.com",
		Status:       "active",
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		CreatedBy:    456,
		AvatarURL:    avatarURL,
	}

	result := DomainToGeneratedEmployee(domainEmp)

	if result == nil {
		t.Fatal("Expected non-nil result for valid input")
	}

	// Check required fields
	if result.Id != 123 {
		t.Errorf("Expected Id=123, got %d", result.Id)
	}
	if result.FirstName != "Иван" {
		t.Errorf("Expected FirstName='Иван', got '%s'", result.FirstName)
	}
	if result.LastName != "Петров" {
		t.Errorf("Expected LastName='Петров', got '%s'", result.LastName)
	}
	if result.Phone != "+79161234567" {
		t.Errorf("Expected Phone='+79161234567', got '%s'", result.Phone)
	}
	if result.Email != "ivan.petrov@example.com" {
		t.Errorf("Expected Email='ivan.petrov@example.com', got '%s'", result.Email)
	}
	if result.Role != "regular" {
		t.Errorf("Expected Role='regular', got '%s'", result.Role)
	}
	if result.Status != "active" {
		t.Errorf("Expected Status='active', got '%s'", result.Status)
	}

	// Check optional fields with values
	if result.MiddleName == nil || *result.MiddleName != "Иванович" {
		t.Errorf("Expected MiddleName='Иванович', got %v", result.MiddleName)
	}
	if result.Position == nil || *result.Position != "Кассир" {
		t.Errorf("Expected Position='Кассир', got %v", result.Position)
	}
	if result.BirthDate == nil || *result.BirthDate != "1990-05-20" {
		t.Errorf("Expected BirthDate='1990-05-20', got %v", result.BirthDate)
	}
	if result.HiredAt == nil || !result.HiredAt.Equal(hiredAt) {
		t.Errorf("Expected HiredAt=%v, got %v", hiredAt, result.HiredAt)
	}
	if result.FiredAt == nil || !result.FiredAt.Equal(firedAt) {
		t.Errorf("Expected FiredAt=%v, got %v", firedAt, result.FiredAt)
	}
	if result.CreatedBy == nil || *result.CreatedBy != 456 {
		t.Errorf("Expected CreatedBy=456, got %v", result.CreatedBy)
	}
	if result.AvatarUrl == nil || *result.AvatarUrl != avatarURL {
		t.Errorf("Expected AvatarUrl='%s', got %v", avatarURL, result.AvatarUrl)
	}

	// Check time fields
	if !result.AddedAt.Equal(addedAt) {
		t.Errorf("Expected AddedAt=%v, got %v", addedAt, result.AddedAt)
	}
	if !result.CreatedAt.Equal(createdAt) {
		t.Errorf("Expected CreatedAt=%v, got %v", createdAt, result.CreatedAt)
	}
	if !result.UpdatedAt.Equal(updatedAt) {
		t.Errorf("Expected UpdatedAt=%v, got %v", updatedAt, result.UpdatedAt)
	}
}

func TestDomainToGeneratedEmployee_MinimalEmployee(t *testing.T) {
	addedAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	updatedAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)

	domainEmp := &domain.Employee{
		ID:           789,
		FirstName:    "Мария",
		LastName:     "Сидорова",
		MiddleName:   nil,
		PasswordHash: "$2b$12$hash",
		Phone:        "+79169876543",
		Position:     nil,
		Role:         "admin",
		BirthDate:    nil,
		HiredAt:      nil,
		FiredAt:      nil,
		AddedAt:      addedAt,
		Email:        "maria@example.com",
		Status:       "onboarding",
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		CreatedBy:    0,  // Zero value
		AvatarURL:    "", // Empty string
	}

	result := DomainToGeneratedEmployee(domainEmp)

	if result == nil {
		t.Fatal("Expected non-nil result for valid input")
	}

	// Check required fields
	if result.Id != 789 {
		t.Errorf("Expected Id=789, got %d", result.Id)
	}
	if result.FirstName != "Мария" {
		t.Errorf("Expected FirstName='Мария', got '%s'", result.FirstName)
	}
	if result.LastName != "Сидорова" {
		t.Errorf("Expected LastName='Сидорова', got '%s'", result.LastName)
	}
	if result.Phone != "+79169876543" {
		t.Errorf("Expected Phone='+79169876543', got '%s'", result.Phone)
	}
	if result.Email != "maria@example.com" {
		t.Errorf("Expected Email='maria@example.com', got '%s'", result.Email)
	}

	// Check optional fields are nil when not set
	if result.MiddleName != nil {
		t.Errorf("Expected MiddleName=nil, got %v", result.MiddleName)
	}
	if result.Position != nil {
		t.Errorf("Expected Position=nil, got %v", result.Position)
	}
	if result.BirthDate != nil {
		t.Errorf("Expected BirthDate=nil, got %v", result.BirthDate)
	}
	if result.HiredAt != nil {
		t.Errorf("Expected HiredAt=nil, got %v", result.HiredAt)
	}
	if result.FiredAt != nil {
		t.Errorf("Expected FiredAt=nil, got %v", result.FiredAt)
	}
	if result.CreatedBy != nil {
		t.Errorf("Expected CreatedBy=nil for zero value, got %v", result.CreatedBy)
	}
	if result.AvatarUrl != nil {
		t.Errorf("Expected AvatarUrl=nil for empty string, got %v", result.AvatarUrl)
	}
}

func TestDomainToGeneratedEmployee_EmptyStringFields(t *testing.T) {
	addedAt := time.Now()
	createdAt := time.Now()
	updatedAt := time.Now()

	domainEmp := &domain.Employee{
		ID:        999,
		FirstName: "Тест",
		LastName:  "Тестов",
		Phone:     "", // Empty phone
		Email:     "", // Empty email
		Role:      "owner",
		Status:    "vacation",
		AddedAt:   addedAt,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		CreatedBy: 0,
		AvatarURL: "", // Empty AvatarURL should convert to nil
	}

	result := DomainToGeneratedEmployee(domainEmp)

	if result == nil {
		t.Fatal("Expected non-nil result")
	}

	// Empty strings should still be passed as empty strings for Phone and Email
	if result.Phone != "" {
		t.Errorf("Expected Phone='', got '%s'", result.Phone)
	}
	if result.Email != "" {
		t.Errorf("Expected Email='', got '%s'", result.Email)
	}

	// Empty AvatarURL should convert to nil
	if result.AvatarUrl != nil {
		t.Errorf("Expected AvatarUrl=nil for empty string, got %v", result.AvatarUrl)
	}
}

func TestDomainToGeneratedEmployee_BirthDateFormatting(t *testing.T) {
	birthDate := time.Date(1985, 12, 25, 14, 30, 0, 0, time.UTC) // Time component should be ignored

	domainEmp := &domain.Employee{
		ID:        1,
		FirstName: "Test",
		LastName:  "User",
		Phone:     "+1234567890",
		Email:     "test@example.com",
		Role:      "regular",
		Status:    "active",
		BirthDate: &birthDate,
		AddedAt:   time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := DomainToGeneratedEmployee(domainEmp)

	if result.BirthDate == nil {
		t.Fatal("Expected BirthDate to be set")
	}

	expectedDate := "1985-12-25"
	if *result.BirthDate != expectedDate {
		t.Errorf("Expected BirthDate='%s', got '%s'", expectedDate, *result.BirthDate)
	}
}

func TestDomainToGeneratedEmployee_TypeConversions(t *testing.T) {
	domainEmp := &domain.Employee{
		ID:        1,
		FirstName: "Test",
		LastName:  "User",
		Phone:     "+1234567890",
		Email:     "test@example.com",
		Role:      "regular",
		Status:    "active",
		AddedAt:   time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: 999, // uint64 -> should convert to *int64
	}

	result := DomainToGeneratedEmployee(domainEmp)

	if result.CreatedBy == nil {
		t.Fatal("Expected CreatedBy to be set")
	}

	if *result.CreatedBy != 999 {
		t.Errorf("Expected CreatedBy=999, got %d", *result.CreatedBy)
	}

	// Verify ID conversion (uint64 -> int64)
	if result.Id != 1 {
		t.Errorf("Expected Id=1, got %d", result.Id)
	}
}
