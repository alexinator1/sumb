package employee

import "time"

type Employee struct {
	ID           uint64     `gorm:"primaryKey;autoIncrement;column:id"`
	FirstName    string     `gorm:"size:100;not null;column:first_name"`
	LastName     string     `gorm:"size:100;not null;column:last_name"`
	MiddleName   *string    `gorm:"size:100;column:middle_name"`
	PasswordHash string     `gorm:"size:255;not null;column:password_hash"`
	PasswordSalt *string    `gorm:"size:255;column:password_salt"`
	Phone        *string    `gorm:"size:30;column:phone"`
	Position     *string    `gorm:"size:100;column:position"`
	Role         string     `gorm:"type:employee_role;not null;default:regular;column:role"`
	BirthDate    *time.Time `gorm:"type:date;column:birth_date"`
	HiredAt      *time.Time `gorm:"column:hired_at"`
	FiredAt      *time.Time `gorm:"column:fired_at"`
	AddedAt      time.Time  `gorm:"column:added_at;not null;default:now()"`
	Email        string     `gorm:"size:255;column:email"`
	Status       string     `gorm:"type:employee_status;not null;default:onboarding;column:status"`
	CreatedAt    time.Time  `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime;column:updated_at"`
	CreatedBy    uint64     `gorm:"column:created_by"`
	AvatarURL    string     `gorm:"size:255;column:avatar_url"`

	// relation to creator (optional)
	Creator *Employee `gorm:"foreignKey:CreatedBy;references:ID;constraint:OnDelete:SET NULL"`
}

func (e *Employee) FullName() string {
	if e == nil {
		return ""
	}
	if e.MiddleName != nil && *e.MiddleName != "" {
		return e.FirstName + " " + *e.MiddleName + " " + e.LastName
	}
	return e.FirstName + " " + e.LastName
}

// IsActive reports whether the employee is not fired (FiredAt == nil).
func (e *Employee) IsActive() bool {
	if e == nil {
		return false
	}
	return e.FiredAt == nil
}
