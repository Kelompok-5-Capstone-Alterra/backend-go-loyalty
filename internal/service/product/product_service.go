package productService

import (
	"backend-go-loyalty/internal/dto"
	"backend-go-loyalty/internal/entity"
	productRepository "backend-go-loyalty/internal/repository/product"
	"context"
)

type IProductService interface {
	InsertProduct(ctx context.Context, req dto.ProductRequest) error
	GetAll(ctx context.Context) (dto.ProductsResponse, error)
	GetProductByID(ctx context.Context, productID uint64) (dto.ProductResponse, error)
	GetProductByCategoryID(ctx context.Context, id uint64) (dto.ProductResponse, error)
	UpdateProduct(ctx context.Context, req dto.ProductUpdateRequest, id uint64) error
	DeleteProduct(ctx context.Context, productID uint64) error
}

type productServiceImpl struct {
	pr productRepository.IProductRepository
}

func NewProductService(pr productRepository.IProductRepository) productServiceImpl {
	return productServiceImpl{
		pr: pr,
	}
}

func (ps productServiceImpl) GetProductByCategoryID(ctx context.Context, id uint64) (dto.ProductResponse, error) {
	product, err := ps.pr.GetProductByCategoryID(ctx, id)
	if err != nil {
		return dto.ProductResponse{}, err
	}
	productResponse := dto.ProductResponse{
		ID:                 product.ID,
		Name:               product.Name,
		Description:        product.Description,
		Provider:           product.Provider,
		ActivePeriod:       product.ActivePeriod,
		Price:              product.Price,
		MinimumTransaction: product.MinimumTransaction,
		Coins:              product.Coins,
		CreatedAt:          product.CreatedAt,
		UpdatedAt:          product.UpdatedAt,
		DeletedAt:          product.DeletedAt,
		Category: dto.CategoryResponse{
			ID:        product.Category.ID,
			Name:      product.Category.Name,
			CreatedAt: product.Category.CreatedAt,
			UpdatedAt: product.Category.UpdatedAt,
			DeletedAt: product.Category.DeletedAt,
		},
	}
	return productResponse, nil
}

func (ps productServiceImpl) GetAll(ctx context.Context) (dto.ProductsResponse, error) {
	products, err := ps.pr.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var productResponses dto.ProductsResponse
	for _, product := range *products {
		item := dto.ProductResponse{
			ID:                 product.ID,
			Name:               product.Name,
			Description:        product.Description,
			Provider:           product.Provider,
			ActivePeriod:       product.ActivePeriod,
			Price:              product.Price,
			MinimumTransaction: product.MinimumTransaction,
			Coins:              product.Coins,
			CreatedAt:          product.CreatedAt,
			UpdatedAt:          product.UpdatedAt,
			DeletedAt:          product.DeletedAt,
			Category: dto.CategoryResponse{
				ID:        product.Category.ID,
				Name:      product.Category.Name,
				CreatedAt: product.Category.CreatedAt,
				UpdatedAt: product.Category.UpdatedAt,
				DeletedAt: product.Category.DeletedAt,
			},
		}
		productResponses = append(productResponses, item)
	}
	if productResponses == nil {
		productResponses = dto.ProductsResponse{}
	}
	return productResponses, nil
}

func (ps productServiceImpl) GetProductByID(ctx context.Context, productID uint64) (dto.ProductResponse, error) {
	product, err := ps.pr.GetProductByID(ctx, productID)
	if err != nil {
		return dto.ProductResponse{}, err
	}
	productResponse := dto.ProductResponse{
		ID:                 product.ID,
		Name:               product.Name,
		Description:        product.Description,
		Provider:           product.Provider,
		ActivePeriod:       product.ActivePeriod,
		Price:              product.Price,
		MinimumTransaction: product.MinimumTransaction,
		Coins:              product.Coins,
		CreatedAt:          product.CreatedAt,
		UpdatedAt:          product.UpdatedAt,
		DeletedAt:          product.DeletedAt,
		Category: dto.CategoryResponse{
			ID:        product.Category.ID,
			Name:      product.Category.Name,
			CreatedAt: product.Category.CreatedAt,
			UpdatedAt: product.Category.UpdatedAt,
			DeletedAt: product.Category.DeletedAt,
		},
	}
	return productResponse, nil
}

func (ps productServiceImpl) InsertProduct(ctx context.Context, req dto.ProductRequest) error {
	product := entity.Product{
		Name:               req.Name,
		Description:        req.Description,
		Provider:           req.Provider,
		ActivePeriod:       req.ActivePeriod,
		Price:              req.Price,
		CategoryID:         req.CategoryID,
		MinimumTransaction: req.MinimumTransaction,
		Coins:              req.Coins,
	}
	err := ps.pr.InsertProduct(ctx, &product)
	if err != nil {
		return err
	}
	return nil
}

func (ps productServiceImpl) UpdateProduct(ctx context.Context, req dto.ProductUpdateRequest, id uint64) error {
	product := entity.Product{
		Name:               req.Name,
		Description:        req.Description,
		Provider:           req.Provider,
		ActivePeriod:       req.ActivePeriod,
		Price:              req.Price,
		CategoryID:         req.CategoryID,
		MinimumTransaction: req.MinimumTransaction,
		Coins:              req.Coins,
	}
	err := ps.pr.UpdateProduct(ctx, &product, id)
	if err != nil {
		return err
	}
	return nil
}

func (ps productServiceImpl) DeleteProduct(ctx context.Context, productID uint64) error {
	err := ps.pr.DeleteProduct(ctx, productID)
	if err != nil {
		return err
	}
	return nil
}
