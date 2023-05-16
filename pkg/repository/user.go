package repository

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	interfaces "github.com/nadeem1815/premium-watch/pkg/repository/interface"
	"github.com/nadeem1815/premium-watch/pkg/utils/idgenerator"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

// UserRegister

func (c *userDatabase) UserRegister(ctx context.Context, user model.UsarDataInput) (model.UserDataOutput, error) {
	var userData model.UserDataOutput
	// query for creating a new entry in users table
	users := domain.Users{
		ID: idgenerator.GenerateID(),
	}
	createUserQuery := `INSERT INTO users(id,name,surname,email_id,password,phone,created_at)
				  VALUES($1,$2,$3,$4,$5,$6,NOW())
				  RETURNING id,name,surname,email_id,phone`

	err := c.DB.Raw(createUserQuery, users.ID, user.Name, user.Surname, user.EmailId, user.Password, user.Phone).Scan(&userData).Error

	if err == nil {
		//query for creating a new entry in the user_infos table.
		insertUserInfoQuery := `INSERT INTO user_infos (is_blocked,users_id) 
								VALUES('f',$1);`
		err = c.DB.Exec(insertUserInfoQuery, userData.ID).Error
	}

	return userData, err
}

func (c *userDatabase) FindByEmail(ctx context.Context, email string) (model.UserLoginVarifier, error) {
	var userData model.UserLoginVarifier
	findUserQuery := `	SELECT 
							u.id,u.name,u.surname,u.email_id,u.password,u.phone,info.is_blocked
						FROM 
							users as u
						FULL OUTER JOIN 
							user_infos as info
						ON 
							u.id=info.id
						WHERE 
							u.email_id=$1;`
	err := c.DB.Raw(findUserQuery, email).Scan(&userData).Error
	return userData, err

}

func (c *userDatabase) BlockUser(ctx context.Context, user_id int) (domain.UserInfo, error) {
	var userInfo domain.UserInfo
	blockQuery := `UPDATE
					user_infos
					SET 
						is_blocked='true',blocked_at=NOW()
					WHERE 
						users_id=$1
					RETURNING *;`
	err := c.DB.Raw(blockQuery, user_id).Scan(&userInfo).Error

	return userInfo, err

}
func (c *userDatabase) UnBlockUser(ctx context.Context, user_id int) (domain.UserInfo, error) {
	var userInfo domain.UserInfo
	unBlockQuery := `UPDATE
					user_infos
					SET
						is_blocked='false'
					WHERE
						users_id=$1
					RETURNING *;`
	err := c.DB.Raw(unBlockQuery, user_id).Scan(&userInfo).Error
	return userInfo, err

}
