package service

import (
	"context"
	"gorm.io/gorm"
	"errors"
	"backend-go-loyalty/internal/dto"
	"backend-go-loyalty/internal/repository/product"
)

type ProductService interface {
	InsertProduct(ctx context.Context, name string, price int, description string, point int) error
	GetAll(ctx context.Context) (dto.ProductResponse, error)
	GetProductByID(ctx context.Context, productID int) (dto.ProductResponse, error)
	UpdateProduct(ctx context.Context) error
	DeleteProduct(ctx context.Context, productID int) error
}

type productServiceImpl struct {
	pr repository.IProductRepository
	db *gorm.DB
	dto dto.ProductResponse
}

func ProvideProductService(pr repository.IProductRepository, db *gorm.DB, dto dto.ProductResponse) *productServiceImpl {
	return &productServiceImpl{
		db: db,
		pr: pr,
		dto: dto,
	}
}

func (ps productServiceImpl) GetAll(ctx context.Context) (dto dto.ProductResponse, error) {
	products, err := ps.pr.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	productResponse := dto.ProductResponse{}
	for _, product := range products {
		var item dto.ProductResponse
		item.ProductID = product.ProductID
		item.Name = product.Name
		item.Price = product.Price
		item.Description = product.Description
		item.Point = product.Point
		productResponse = append(productResponse, &item)
	}
	return productResponse, nil
}

func (ps productServiceImpl) GetProductByID(ctx context.Context, productID int) (dto dto.ProductResponse, error) {
	prod, err := ps.pr.GetProductByID(ctx, productID)
	if err != nil {
		return dto.ProductResponse{}, err
	}
	if len(prod) < 1 {
		return nil, errors.New("no product found")
	}
	productResponse := dto.ProductResponse{}
	for _, product := range prod {
		var item dto.ProductResponse
		item.ProductID = product.ProductID
		item.Name = product.Name
		item.Price = product.Price
		item.Description = product.Description
		item.Point = product.Point
	}
	return productResponse, nil
}

func (ps productServiceImpl) InsertProduct(ctx context.Context, name string, price int, description string, point int) error {
	err := ps.pr.InsertProduct(ctx, product.Name, product.Price, product.Description, product.Price)
	if err != nil {
		return 0, err
	}
	return nil
}

func (ps productServiceImpl) UpdateProduct(ctx context.Context) error {
	err := ps.pr.UpdateProduct(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ps productServiceImpl) DeleteProduct(ctx context.Context, productID int) error {
	err := ps.pr.DeleteProduct(ctx, productID)
	if err != nil {
		return err
	}
	return nil
}