package interfaces

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
)

type UserRepository interface {
	UserRegister(ctx context.Context, user model.UsarDataInput) (model.UserDataOutput, error)
	FindByEmail(ctx context.Context, email string) (model.UserLoginVarifier, error)
	BlockUser(ctx context.Context, user_id int) (domain.UserInfo, error)
	UnBlockUser(ctx context.Context, user_id int) (domain.UserInfo, error)
}
