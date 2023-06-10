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

func (cr *productUseCase) FindCategoryById(ctx context.Context, categoriesid int) (domain.ProductCategory, error) {
	categoryId, err := cr.productRepo.FindCategoryById(ctx, categoriesid)
	return categoryId, err
}

func (cr *productUseCase) CreateProduct(ctx context.Context, createProduct domain.Product) (domain.Product, error) {
	createProducts, err := cr.productRepo.CreateProduct(ctx, createProduct)
	return createProducts, err
}

func (cr *productUseCase) ListAllProducts() ([]model.OutPutProduct, error) {
	viewAllProduct, err := cr.productRepo.ListAllProducts()
	return viewAllProduct, err
}

func (cr *productUseCase) UpdateProduct(ctx context.Context, updataproduct domain.Product) (domain.Product, error) {
	updateProductItem, err := cr.productRepo.UpdateProduct(ctx, updataproduct)
	return updateProductItem, err
}

func (cr *productUseCase) DeleteProduct(ctx context.Context, id int) error {
	err := cr.productRepo.DeleteProduct(ctx, id)
	return err
}

func (cr *productUseCase) CreateCoupon(ctx context.Context, createdCoupon model.CreatCoupon) (domain.Coupon, error) {
	coupon, err := cr.productRepo.CreateCoupon(ctx, createdCoupon)
	if err != nil {

		return domain.Coupon{}, err
	}
	return coupon, nil
}

func (cr *productUseCase) UpdateCoupon(ctx context.Context, couponInfo model.UpdatCoupon) (domain.Coupon, error) {
	updatedCoupon, err := cr.productRepo.UpdateCoupon(ctx, couponInfo)
	if err != nil {
		return domain.Coupon{}, err

	}
	// if updatedCoupon.ID == 0 {

	// 	return domain.Coupon{}, fmt.Errorf("failed coupo updating")
	// }
	return updatedCoupon, nil
}
func (cr *productUseCase) DeleteCoupon(ctx context.Context, couponID int) error {
	err := cr.productRepo.DeleteCoupon(ctx, couponID)
	return err
}
