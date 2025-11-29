package employee

import (
	"context"
	"errors"

	entity "github.com/alexinator1/sumb/back/internal/domain/employee/entity"
	"gorm.io/gorm"
)

// EmployeeRepo репозиторий для пользователей
type EmployeeRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *EmployeeRepo {
	return &EmployeeRepo{db: db}
}

func (r *EmployeeRepo) GetByID(ctx context.Context, id uint64) (*entity.Employee, error) {
	var e entity.Employee
	if err := r.db.WithContext(ctx).First(&e, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}
	return &e, nil
}

func (r *EmployeeRepo) Create(ctx context.Context, e *entity.Employee) error {
	return r.CreateWithDB(ctx, r.db, e)
}

func (r *EmployeeRepo) CreateWithDB(ctx context.Context, db *gorm.DB, e *entity.Employee) error {
	if err := db.WithContext(ctx).Create(e).Error; err != nil {
		return err
	}
	return nil
}
