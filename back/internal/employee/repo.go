package employee

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

// Repo репозиторий для пользователей
type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) GetByID(ctx context.Context, id uint64) (*Employee, error) {
	var e Employee
	if err := r.db.WithContext(ctx).First(&e, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &e, nil
}
