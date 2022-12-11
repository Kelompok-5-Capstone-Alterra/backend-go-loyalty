package categoryRepository

import (
	"backend-go-loyalty/internal/entity"
	"backend-go-loyalty/internal/model"
	"context"

	"gorm.io/gorm"
)

type ICategoryRepository interface {
	GetAllCategory(ctx context.Context) (entity.Categories, error)
	GetCategoryByID(ctx context.Context, id uint64) (entity.Category, error)
	CreateCategory(ctx context.Context, req entity.Category) error
	UpdateCategory(ctx context.Context, req entity.Category, id uint64) error
	DeleteCategory(ctx context.Context, id uint64) error
}

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) categoryRepository {
	return categoryRepository{
		DB: db,
	}
}

func (cr categoryRepository) GetAllCategory(ctx context.Context) (entity.Categories, error) {
	categories := entity.Categories{}
	err := cr.DB.Find(&categories).Error
	return categories, err
}
func (cr categoryRepository) GetCategoryByID(ctx context.Context, id uint64) (entity.Category, error) {
	category := entity.Category{}
	err := cr.DB.First(&category, id).Error
	return category, err
}
func (cr categoryRepository) CreateCategory(ctx context.Context, req entity.Category) error {
	err := cr.DB.Create(&req).Error
	return err
}
func (cr categoryRepository) UpdateCategory(ctx context.Context, req entity.Category, id uint64) error {
	err := cr.DB.Model(&model.Category{}).Where("id = ?", id).Updates(req).Error
	return err
}
func (cr categoryRepository) DeleteCategory(ctx context.Context, id uint64) error {
	category := entity.Category{}
	err := cr.DB.Delete(&category, id).Error
	return err
}
