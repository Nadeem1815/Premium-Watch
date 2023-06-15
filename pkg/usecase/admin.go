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

type adminUseCase struct {
	adminRepo interfaces.AdminRepository
}

func NewAdminUseCase(repo interfaces.AdminRepository) services.AdminUseCase {
	return &adminUseCase{
		adminRepo: repo,
	}
}

func (c *adminUseCase) SaveAdmin(ctx context.Context, admin domain.Admin) error {

	// if admins, err := c.adminRepo.FindAdmin(ctx, admin); err != nil {
	// 	return err
	// } else if admins.ID != 0 {
	// 	return errors.New("admin already exist this same details")

	// }
	// generate hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return err
	}
	admin.Password = string(hash)

	err = c.adminRepo.AdminSave(ctx, admin)
	if err != nil {
		return err

	}

	return nil
}

func (c *adminUseCase) AdminLogin(ctx context.Context, input model.AdminLogin) (string, model.AdminDataOutput, error) {
	var details model.AdminDataOutput

	// 1. find the adminDAta given email
	admininfo, err := c.adminRepo.FindAdmin(ctx, input)
	details.ID = admininfo.ID
	details.UserName = admininfo.UserName
	details.Email = admininfo.Email
	if err != nil {
		fmt.Println("1")
		return "", details, fmt.Errorf("error finding admin data")
	}

	// 2. compare and hash the password
	if err := bcrypt.CompareHashAndPassword([]byte(admininfo.Password), []byte(input.Password)); err != nil {

		return "", details, err
	}

	// if correct password genterate the token
	// jwt token creating send the cookie
	claims := jwt.MapClaims{
		"id":  admininfo.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed string

	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", model.AdminDataOutput{}, err
	}

	return ss, details, nil
}

func (c *adminUseCase) ListAllUsers() ([]domain.Users, error) {
	users, err := c.adminRepo.ListAllUsers()
	return users, err

}

func (c *adminUseCase) FindUserId(ctx context.Context, userId string) (domain.Users, error) {
	findUser, err := c.adminRepo.FindUserId(ctx, userId)
	return findUser, err
}

func (c *adminUseCase) DashBoard(ctx context.Context) (model.AdminDashBoard, error) {
	adminDashBoard, err := c.adminRepo.DashBoard(ctx)
	return adminDashBoard, err
}

func (c *adminUseCase) SalesRepo(ctx context.Context) ([]model.SalesReport, error) {
	sales, err := c.adminRepo.SalesRepo(ctx)
	fmt.Println("usecAse:",)
	return sales, err
}
