package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	interfaces "github.com/nadeem1815/premium-watch/pkg/repository/interface"
	"github.com/nadeem1815/premium-watch/pkg/utils/idgenerator"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{DB}
}

func (c *adminDatabase) AdminSave(ctx context.Context, admin domain.Admin) error {

	admins := domain.Admin{
		ID: idgenerator.GenerateID(),
	}
	insertquery := `INSERT INTO admins (id,user_name,email,password)VALUES($1,$2,$3,$4)
			RETURNING id,user_name,email `
	err := c.DB.Raw(insertquery, admins.ID, admin.UserName, admin.Email, admin.Password).Scan(&admin).Error
	return err
}

func (c *adminDatabase) FindAdmin(ctx context.Context, admin model.AdminLogin) (domain.Admin, error) {
	var adminDetail domain.Admin
	query := `SELECT *FROM admins WHERE email=? `
	err := c.DB.Raw(query, admin.Email).Scan(&adminDetail).Error
	if err != nil {
		return adminDetail, errors.New("can't find admin")
	}
	fmt.Println(adminDetail)
	return adminDetail, nil
}

//	func (c *adminDatabase)LoginAdmin(ctx context.Context,email string)(domain.Admin,error){
//		var admiData domain.Admin
//		findAdminQuery:=`SELECT`
//	}
func (c *adminDatabase) ListAllUsers() ([]domain.Users, error) {
	var users []domain.Users
	listUserQuery := `SELECT
					id,name,email_id,phone
				FROM
					users;`
	err := c.DB.Raw(listUserQuery).Scan(&users).Error
	if err != nil {
		return []domain.Users{}, err
	}
	return users, nil

}

func (c *adminDatabase) FindUserId(ctx context.Context, userId string) (domain.Users, error) {
	tx := c.DB.Begin()
	var isExists bool
	if err := tx.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id=$1)", userId).Scan(&isExists).Error; err != nil {
		tx.Rollback()
		return domain.Users{}, err
	}
	if !isExists {
		tx.Rollback()
		return domain.Users{}, fmt.Errorf("no user present")
	}
	var user domain.Users
	findUserIdQuery := `SELECT 
					  			id,name,email_id,phone
						 FROM
								users
						WHERE 
							id=$1;`
	if err := tx.Raw(findUserIdQuery, userId).Scan(&user).Error; err != nil {
		tx.Rollback()
		return domain.Users{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return domain.Users{}, err
	}
	return user, nil

}
