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

	CreateCoupon(ctx context.Context, createdCoupon model.CreatCoupon) (domain.Coupon, error)
	UpdateCoupon(ctx context.Context, couponInfo model.UpdatCoupon) (domain.Coupon, error)
	DeleteCoupon(ctx context.Context, couponID int) error
	ViewAllCoupon() ([]domain.Coupon, error)
	ViewCouponById(ctx context.Context, couponID int) (domain.Coupon, error)
	CouponUsed(ctx context.Context, userID string, couponID int) (bool, error)
}
