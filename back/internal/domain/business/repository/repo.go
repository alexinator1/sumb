package business

import (
	"context"
	"errors"

	entity "github.com/alexinator1/sumb/back/internal/domain/business/entity"
	"gorm.io/gorm"
)

type BusinessRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *BusinessRepo {
	return &BusinessRepo{db: db}
}

func (r *BusinessRepo) GetByID(ctx context.Context, id uint64) (*entity.Business, error) {
	var b entity.Business
	if err := r.db.WithContext(ctx).First(&b, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &b, nil
}

func (r *BusinessRepo) Create(ctx context.Context, b *entity.Business) error {
	if err := r.db.WithContext(ctx).Create(b).Error; err != nil {
		return err
	}
	return nil
}

func (r *BusinessRepo) Update(ctx context.Context, b *entity.Business) error {
	if err := r.db.WithContext(ctx).Save(b).Error; err != nil {
		return err
	}
	return nil
}