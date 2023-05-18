package handler

import (
	services "github.com/nadeem1815/premium-watch/pkg/usecase/interface"
)

type CartHandler struct {
	cartUseCase services.CartUseCase
}


func NewCartHandler(usecase services.CartUseCase)*CartHandler{
	return &CartHandler{
		cartUseCase: usecase,
	}
}

