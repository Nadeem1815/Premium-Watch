package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/nadeem1815/premium-watch/pkg/domain"
	interfaces "github.com/nadeem1815/premium-watch/pkg/repository/interface"
	services "github.com/nadeem1815/premium-watch/pkg/usecase/interface"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) UserRegister(ctx context.Context, input model.UsarDataInput) (model.UserDataOutput, error) {
	// hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return model.UserDataOutput{}, err

	}
	input.Password = string(hash)
	// passing for userdata in repo
	user, err := c.userRepo.UserRegister(ctx, input)
	

	return user, err
}

func (c *userUseCase) LoginWithEmail(ctx context.Context, input model.UserLoginEmail) (string, model.UserDataOutput, error) {
	var userData model.UserDataOutput
	// 1. Get user email id and password from handler

	user, err := c.userRepo.FindByEmail(ctx, input.EmailId)
	if err != nil {
		return "", userData, fmt.Errorf("error finding userData")

	}
	// 2. Find user details from database, using given email
	if input.EmailId == "" {
		return "", userData, fmt.Errorf("no such userdata found")

	}
	fmt.Println("checkpoint 0")

	// 3. Compare given password and password from database

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", userData, err

	}
	if user.IsBlocked {
		return "", userData, fmt.Errorf("account is blocked")
	}

	// 4. If correct password, generate jwt token, and pass it using cookie
	// creating jwt token sending it in cookie
	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// singed string
	ss, err := token.SignedString([]byte("secret"))

	// 5. Return user data and jwt token to handler
	userData.ID, userData.Name, userData.Surname, userData.EmailId, userData.Phone = user.ID, user.Name, user.Surname, user.EmailId, user.Phone
	return ss, userData, err

}
func (c *userUseCase) BlockUser(ctx context.Context, user_id string) (domain.UserInfo, error) {
	blockedUser, err := c.userRepo.BlockUser(ctx, user_id)
	return blockedUser, err

}
func (c *userUseCase) UnBlockUser(ctx context.Context, user_id string) (domain.UserInfo, error) {
	unBlockedUser, err := c.userRepo.UnBlockUser(ctx, user_id)
	return unBlockedUser, err
}
