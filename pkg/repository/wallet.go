package repository

import (
	"context"
	"fmt"

	"github.com/nadeem1815/premium-watch/pkg/domain"
)

// find wallet for user
func (cr *OrderDatabase) UserWallet(ctx context.Context, userID string) (domain.Wallet, error) {
	var wallet domain.Wallet

	findWalletQuery := `SELECT *FROM wallets WHERE user_id=$1;`

	err := cr.DB.Raw(findWalletQuery, userID).Scan(&wallet).Error

	if err != nil {
		return domain.Wallet{}, fmt.Errorf("failed to find wallet where userID")

	}
	return wallet, nil
}

// create a newWallet for user

func (cr *OrderDatabase) SaveWallet(ctx context.Context, userID string) (domain.Wallet, error) {
	var wallet domain.Wallet

	query := `INSERT INTO wallets (user_id,wallet_balance) VALUES ($1,$2);`

	err := cr.DB.Raw(query, userID, 0).Scan(&wallet).Error
	if err != nil {
		return domain.Wallet{}, err
	}

	return wallet, nil
}
