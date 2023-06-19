package usecase

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
)

func (cr *OrderUseCase) UserWallet(ctx context.Context, userID string) (domain.Wallet, error) {

	// find user wallet
	wallet, err := cr.orderRepo.UserWallet(ctx, userID)
	if err != nil {
		return domain.Wallet{}, err

	} else if wallet.ID == 0 {
		wallet, err = cr.orderRepo.SaveWallet(ctx, userID)
		if err != nil {
			return domain.Wallet{}, err
		}

	}
	return wallet, nil
}
