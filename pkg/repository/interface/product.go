package interfaces

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
)

type ProductRepository interface {
	CreateCategory(ctx context.Context, categoryName string) (domain.ProductCategory, error)
	ViewAllCategory() ([]domain.ProductCategory, error)
	FindCategoryById(ctx context.Context, categoriesid int) (domain.ProductCategory, error)

	CreateProduct(ctx context.Context, createProduct domain.Product) (domain.Product, error)
	ListAllProducts() ([]model.OutPutProduct, error)
	UpdateProduct(ctx context.Context, updateProduct domain.Product) (domain.Product, error)
	DeleteProduct(ctx context.Context, id int) error

	CreateCoupon(ctx context.Context, createdCoupon model.CreateCoupon) (domain.Coupon, error)
}
