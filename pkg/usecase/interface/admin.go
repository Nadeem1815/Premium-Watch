package interfaces

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
)

type AdminUseCase interface {
	SaveAdmin(ctx context.Context, admin domain.Admin) error
	AdminLogin(ctx context.Context, input model.AdminLogin) (string, model.AdminDataOutput, error)
	ListAllUsers() ([]domain.Users, error)
	FindUserId(ctx context.Context, userId int) (domain.Users, error)
}
