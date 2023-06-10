package interfaces

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
)

type ProductUseCase interface {
	CreateCategory(ctx context.Context, CategoryName string) (domain.ProductCategory, error)
	ViewAllCategory() ([]domain.ProductCategory, error)
	FindCategoryById(ctx context.Context, categoriesid int) (domain.ProductCategory, error)

	CreateProduct(ctx context.Context, createProduct domain.Product) (domain.Product, error)
	ListAllProducts() ([]model.OutPutProduct, error)
	UpdateProduct(ctx context.Context, updataproduct domain.Product) (domain.Product, error)
	DeleteProduct(ctx context.Context, id int) error

	CreateCoupon(ctx context.Context, createdCoupon model.CreateCoupon) (domain.Coupon, error)
	UpdateCoupon(ctx context.Context, UpdatedCoupon model.UpdatCoupon) (domain.Coupon, error)
}
