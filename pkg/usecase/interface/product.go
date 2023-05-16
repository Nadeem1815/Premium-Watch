package interfaces

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
)

type ProductUseCase interface {
	CreateCategory(ctx context.Context, CategoryName string) (domain.ProductCategory, error)
	ViewAllCategory() ([]domain.ProductCategory, error)
	CreateProduct(ctx context.Context, createProduct domain.Product) (domain.Product, error)
	ListAllProducts() ([]model.OutPutProduct, error)
}
