package db

import (
	"fmt"

	config "github.com/nadeem1815/premium-watch/pkg/config"
	domain "github.com/nadeem1815/premium-watch/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if dbErr != nil {
		return db, dbErr
	}
	dbErr = db.AutoMigrate(

		// user tables
		&domain.Users{},
		&domain.UserInfo{},

		// admin table
		&domain.Admin{},

		// product table
		&domain.ProductCategory{},
		&domain.Product{},

		//cart table
		&domain.Cart{},
		&domain.CartItems{},

		// address table
		&domain.Address{},

		// order table
		&domain.Order{},
		&domain.OrderItem{},
		&domain.OrderStatus{},
		&domain.DeliveryStatus{},
		&domain.Return{},

		// payment table
		&domain.PaymentDetails{},
		&domain.PaymentStatus{},
		&domain.PaymentMethod{},

		// coupon table
		&domain.Coupon{},
	)

	return db, dbErr
}
