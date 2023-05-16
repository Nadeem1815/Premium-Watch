package usecase

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	interfaces "github.com/nadeem1815/premium-watch/pkg/repository/interface"
	services "github.com/nadeem1815/premium-watch/pkg/usecase/interface"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
)

type productUseCase struct {
	productRepo interfaces.ProductRepository
}

func NewProductUseCase(repo interfaces.ProductRepository) services.ProductUseCase {
	return &productUseCase{
		productRepo: repo,
	}
}

func (cr *productUseCase) CreateCategory(ctx context.Context, CategoryName string) (domain.ProductCategory, error) {
	createdCategory, err := cr.productRepo.CreateCategory(ctx, CategoryName)
	return createdCategory, err
}

func (cr *productUseCase) ViewAllCategory() ([]domain.ProductCategory, error) {
	viewAllCategories, err := cr.productRepo.ViewAllCategory()
	return viewAllCategories, err
}

func (cr *productUseCase) CreateProduct(ctx context.Context, createProduct domain.Product) (domain.Product, error) {
	createProducts, err := cr.productRepo.CreateProduct(ctx, createProduct)
	return createProducts, err
}

func (cr *productUseCase) ListAllProducts() ([]model.OutPutProduct, error) {
	viewAllProduct, err := cr.productRepo.ListAllProducts()
	return viewAllProduct, err
}
