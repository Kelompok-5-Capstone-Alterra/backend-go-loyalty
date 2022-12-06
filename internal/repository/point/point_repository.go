package pointRepository

import (
	"backend-go-loyalty/internal/entity"
	"backend-go-loyalty/internal/model"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IPointRepository interface {
	GetAllPoints(ctx context.Context) (entity.UserCoins, error)
	GetPoint(ctx context.Context, id uuid.UUID) (entity.UserCoin, error)
	UpdatePoint(ctx context.Context, id uuid.UUID, req entity.UserCoin) error
}

type pointRepository struct {
	DB *gorm.DB
}

func NewPointRepository(db *gorm.DB) pointRepository {
	return pointRepository{
		DB: db,
	}
}

func (pr pointRepository) GetAllPoints(ctx context.Context) (entity.UserCoins, error) {
	points := entity.UserCoins{}
	err := pr.DB.Model(&model.UserCoin{}).Preload("User").Find(&points).Error
	return points, err
}

func (pr pointRepository) GetPoint(ctx context.Context, id uuid.UUID) (entity.UserCoin, error) {
	point := entity.UserCoin{}
	err := pr.DB.Model(&model.UserCoin{}).Preload("User").Where("user_id = ?", id).First(&point).Error
	return point, err
}
func (pr pointRepository) UpdatePoint(ctx context.Context, id uuid.UUID, req entity.UserCoin) error {
	err := pr.DB.Model(&model.UserCoin{}).Where("user_id = ?", id).Updates(req).Error
	return err
}
