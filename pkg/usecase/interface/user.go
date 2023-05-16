package interfaces

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
)

type UserUseCase interface {
	UserRegister(ctx context.Context, input model.UsarDataInput) (model.UserDataOutput, error)
	LoginWithEmail(ctx context.Context, input model.UserLoginEmail) (string, model.UserDataOutput, error)
	BlockUser(ctx context.Context, user_id int) (domain.UserInfo, error)
	UnBlockUser(ctx context.Context,user_id int)(domain.UserInfo,error)
}
