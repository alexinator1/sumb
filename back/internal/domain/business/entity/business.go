package entity

import "time"

type Business struct {
	ID              uint64     `gorm:"primaryKey;autoIncrement;column:id"`
	Name            string     `gorm:"size:200;not null;column:name"`
	Description     *string    `gorm:"column:description"`
	OwnerFirstName  string     `gorm:"size:100;column:owner_first_name"`
	OwnerLastName   string     `gorm:"size:100;column:owner_last_name"`
	OwnerMiddleName *string    `gorm:"size:100;column:owner_middle_name"`
	OwnerEmail      string     `gorm:"size:255;column:owner_email"`
	OwnerPhone      string     `gorm:"size:30;column:owner_phone"`
	LogoID          *string    `gorm:"size:255;column:logo_id"`
	CreatedAt       time.Time  `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime;column:updated_at"`
	IsWorking       bool       `gorm:"column:is_working"`
	DeletedAt       *time.Time `gorm:"column:deleted_at"`
	OwnerID         *uint64    `gorm:"column:owner_id"`
}

// TableName specifies the table name for GORM
func (Business) TableName() string {
	return "business"
}	
