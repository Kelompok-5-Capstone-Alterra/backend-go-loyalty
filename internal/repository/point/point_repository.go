package pointRepository

import (
	"backend-go-loyalty/internal/entity"
	"backend-go-loyalty/internal/model"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IPointRepository interface {
	GetAllPoints(ctx context.Context) (entity.Points, error)
	GetPoint(ctx context.Context, id uuid.UUID) (entity.Point, error)
	UpdatePoint(ctx context.Context, id uuid.UUID, req entity.Point) error
}

type pointRepository struct {
	DB *gorm.DB
}

func NewPointRepository(db *gorm.DB) pointRepository {
	return pointRepository{
		DB: db,
	}
}

func (pr pointRepository) GetAllPoints(ctx context.Context) (entity.Points, error) {
	points := entity.Points{}
	err := pr.DB.Model(&model.Point{}).Preload("UserID").Preload("RoleID").Find(&points).Error
	return points, err
}

func (pr pointRepository) GetPoint(ctx context.Context, id uuid.UUID) (entity.Point, error) {
	point := entity.Point{}
	err := pr.DB.Model(&model.Point{}).Preload("UserID").Preload("RoleID").First(&point, id).Error
	return point, err
}
func (pr pointRepository) UpdatePoint(ctx context.Context, id uuid.UUID, req entity.Point) error {
	err := pr.DB.Model(&model.Point{}).Where("id = ?", id).Updates(req).Error
	return err
}
