//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/nadeem1815/premium-watch/pkg/api"
	handler "github.com/nadeem1815/premium-watch/pkg/api/handler"
	config "github.com/nadeem1815/premium-watch/pkg/config"
	db "github.com/nadeem1815/premium-watch/pkg/db"
	repository "github.com/nadeem1815/premium-watch/pkg/repository"
	usecase "github.com/nadeem1815/premium-watch/pkg/usecase"
)

func InitializerAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		// connect to database

		db.ConnectDatabase,

		// handler
		handler.NewAdminHandler,
		handler.NewUserHandler,
		handler.NewProductHandler,
		handler.NewCartHandler,
		handler.NewOrderHandler,

		// usecase
		usecase.NewAdminUseCase,
		usecase.NewUserUseCase,
		usecase.NewProductUseCase,
		usecase.NewCartUseCase,
		usecase.NewOrderUseCase,

		// repository
		repository.NewAdminRepository,
		repository.NewUserRepository,
		repository.NewProductRepository,
		repository.NewCartRepository,
		repository.NewOrderRepository,

		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
