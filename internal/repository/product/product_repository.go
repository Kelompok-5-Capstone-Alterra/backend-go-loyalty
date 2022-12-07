package productRepository

import (
	"backend-go-loyalty/internal/entity"
	"backend-go-loyalty/internal/model"
	"context"

	"gorm.io/gorm"
)

type IProductRepository interface {
	InsertProduct(ctx context.Context, product *entity.Product) error
	GetAll(ctx context.Context) (*entity.Products, error)
	GetProductByID(ctx context.Context, id uint64) (*entity.Product, error)
	UpdateProduct(ctx context.Context, p *entity.Product, id uint64) error
	DeleteProduct(ctx context.Context, id uint64) error
}
type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &productRepository{db}
}

func (pr *productRepository) GetAll(ctx context.Context) (*entity.Products, error) {
	var products entity.Products
	err := pr.DB.Model(&model.Product{}).Preload("Category").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return &products, nil
}

func (pr *productRepository) GetProductByID(ctx context.Context, id uint64) (*entity.Product, error) {
	var product entity.Product
	err := pr.DB.Model(&model.Product{}).Preload("Category").First(&product, id).Error

	if err != nil {
		return nil, err
	}
	return &product, err
}

func (pr *productRepository) InsertProduct(ctx context.Context, product *entity.Product) error {
	err := pr.DB.Create(&product).Error

	if err != nil {
		return err
	}
	return nil
}

func (pr *productRepository) UpdateProduct(ctx context.Context, p *entity.Product, id uint64) error {
	err := pr.DB.Model(&model.Product{}).Where("id = ?", id).Updates(p).Error
	if err != nil {
		return err
	}
	return nil
}

func (pr *productRepository) DeleteProduct(ctx context.Context, id uint64) error {
	var product entity.Product
	err := pr.DB.Delete(&product, id).Error

	if err != nil {
		return err
	}

	return nil

}
