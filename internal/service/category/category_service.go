package categoryService

import (
	"backend-go-loyalty/internal/dto"
	"backend-go-loyalty/internal/entity"
	categoryRepository "backend-go-loyalty/internal/repository/category"
	"context"
	"time"
)

type ICategorySerivce interface {
	GetAllCategories(ctx context.Context) (dto.CategoriesResponse, error)
	GetCategoryByID(ctx context.Context, id uint64) (dto.CategoryResponse, error)
	CreateCategory(ctx context.Context, req dto.CategoryRequest) error
	UpdateCategory(ctx context.Context, id uint64, req dto.CategoryRequest) error
	DeleteCategory(ctx context.Context, id uint64) error
}

type categoryService struct {
	cr categoryRepository.ICategoryRepository
}

func NewCategoryService(cr categoryRepository.ICategoryRepository) categoryService {
	return categoryService{
		cr: cr,
	}
}

func (cs categoryService) GetAllCategories(ctx context.Context) (dto.CategoriesResponse, error) {
	data, err := cs.cr.GetAllCategory(ctx)
	if err != nil {
		return nil, err
	}
	var categories dto.CategoriesResponse
	for _, val := range data {
		category := dto.CategoryResponse{
			ID:        val.ID,
			Name:      val.Name,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
			DeletedAt: val.DeletedAt,
		}
		categories = append(categories, category)
	}
	if categories == nil {
		categories = dto.CategoriesResponse{}
	}
	return categories, err
}
func (cs categoryService) GetCategoryByID(ctx context.Context, id uint64) (dto.CategoryResponse, error) {
	data, err := cs.cr.GetCategoryByID(ctx, id)
	if err != nil {
		return dto.CategoryResponse{}, err
	}
	category := dto.CategoryResponse{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}
	return category, nil
}
func (cs categoryService) CreateCategory(ctx context.Context, req dto.CategoryRequest) error {
	category := entity.Category{
		Name:      req.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := cs.cr.CreateCategory(ctx, category)
	return err
}
func (cs categoryService) UpdateCategory(ctx context.Context, id uint64, req dto.CategoryRequest) error {
	category := entity.Category{
		Name: req.Name,
	}
	err := cs.cr.UpdateCategory(ctx, category, id)
	return err
}
func (cs categoryService) DeleteCategory(ctx context.Context, id uint64) error {
	err := cs.cr.DeleteCategory(ctx, id)
	return err
}
