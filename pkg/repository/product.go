package repository

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	interfaces "github.com/nadeem1815/premium-watch/pkg/repository/interface"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
	"gorm.io/gorm"
)

type productDataBase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &productDataBase{DB}
}

func (c *productDataBase) CreateCategory(ctx context.Context, categoryName string) (domain.ProductCategory, error) {
	var createdCategory domain.ProductCategory

	categoryCreatequery := `INSERT INTO product_categories(category_name)
						 VALUES ($1)
						 RETURNING  id,category_name`
	err := c.DB.Raw(categoryCreatequery, categoryName).Scan(&createdCategory).Error
	return createdCategory, err
}

func (c *productDataBase) ViewAllCategory() ([]domain.ProductCategory, error) {
	var allcategories []domain.ProductCategory
	allcategoryquery := `SELECT *FROM product_categories`
	err := c.DB.Raw(allcategoryquery).Scan(&allcategories).Error
	if err != nil {
		return []domain.ProductCategory{}, err
	}
	return allcategories, err

}

func (cr *productDataBase) CreateProduct(ctx context.Context, createProduct domain.Product) (domain.Product, error) {
	var createdProducts domain.Product
	productItemsCreatequery := `INSERT INTO products(product_category_id,name,brand,colour,description,price,stock,product_image,sku,created_at)
							  VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,NOW())
							  RETURNING *`
	err := cr.DB.Raw(productItemsCreatequery,
		createProduct.ProductCategoryID,
		createProduct.Name, createProduct.Brand,
		createProduct.Colour,
		createProduct.Description,
		createProduct.Price,
		createProduct.Stock,
		createProduct.ProductImage,
		createProduct.SKU).Scan(&createdProducts).Error
	return createdProducts, err

}

func (cr *productDataBase) ListAllProducts() ([]model.OutPutProduct, error) {
	var allProduct []model.OutPutProduct
	viewAllProductQuery := `SELECT p.*, c.category_name FROM products p LEFT JOIN product_categories c ON p.product_category_id = c.id;`
	err := cr.DB.Raw(viewAllProductQuery).Scan(&allProduct).Error
	if err != nil {
		return []model.OutPutProduct{}, err
	}
	return allProduct, err
}

func (cr *productDataBase) UpdateProduct(ctx context.Context, updateProduct domain.Product) (domain.Product, error) {
	var updateProductItem domain.Product
	updateProductQuery := `UPDATE products
						 SET             
							 product_category_id=$1, 
							 name=$2,            
							 brand=$3,         
							 colour=$4,           
							 description=$5,      
							 price=$6,            
							 stock=$7,            
							 product_image=$8,     
							 sku=$9,    
							 updated_at=NOW()
						WHERE id=$10
						RETURNING id,product_category_id,name,brand,colour,description,price,stock,product_image,sku,updated_at`
	err := cr.DB.Raw(updateProductQuery,
		updateProduct.ProductCategoryID,
		updateProduct.Name,
		updateProduct.Brand,
		updateProduct.Colour,
		updateProduct.Description,
		updateProduct.Price,
		updateProduct.Stock,
		updateProduct.ProductImage,
		updateProduct.SKU,
		updateProduct.ID).Scan(&updateProductItem).Error
	return updateProductItem, err

}
