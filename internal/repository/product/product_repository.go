package repository

import (
	"backend-go-loyalty/entity"
	"gorm.io/gorm"
)

type IProductRepository interface {
	InsertProduct(product *entity.Product) (*entity.Product, error)
	GetAll(p []*entity.Product) ([]*entity.Product, error)
	GetProductByID(id int) (*entity.Product, error)
	UpdateProduct(p *entity.Product, id int) (*entity.Product, error)
	DeleteProduct(id int) error
}
type productRepository struct {
	DB *gorm.DB
}

type NewProductRepository(db *gorm.DB) IProductRepository {
	return &productRepository{db}
}

func (pr *productRepository) GetAll(p []*entity.Product) ([]*entity.Product, error) {
	err := pr.Db.Find(&p).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (pr *productRepository) GetProductByID(id int) (*entity.Product, error) {
	var product entity.Product
	err := pr.Db.Where("id = ?", id).First(&product).Error
	
	if err != nil {
		return nil, err
	}
	return &product, err
}

func (pr *productRepository) InsertProduct(product *entity.Product) (*entity.Product, error) {
	err := pr.Db.Where("name =? AND price =? AND description =? AND points =?", product.Name, product.Price, product.Description, product.Point).First(&product).Error
	err = pr.Db.Create(&product).Error
	
	if err != nil {
		return nil, err
	}
	return product, err
}

func (pr *productRepository) UpdateProduct(p *entity.Product, id int) (*entity.Product, error) {
	var product entity.Product
	err := pr.Db.Find(&product, id).Error
	product.Name = p.UserName
	product.Price = p.Price
	product.Description = p.Description
	product.Point = p.Point

	err = pr.Db.Save(&product).Error

	if err != nil {
		return nil, err
	}

	return p, err
}

func (pr *productRepository) DeleteProduct(id int) error {
	var product entity.Product
	err := pr.Db.Delete(&product, id).Error

	if err != nil {
		return err
	}

	return nil

}