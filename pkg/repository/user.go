package repository

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	interfaces "github.com/nadeem1815/premium-watch/pkg/repository/interface"
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
	// users := domain.Users{
	// 	ID: idgenerator.GenerateID(),
	// }

	// 	insertQuery := `INSERT INTO users (name,email,mobile,password)VALUES($1,$2,$3,$4)
	// 	RETURNING id,name,email,mobile`
	// err = c.DB.Raw(insertQuery, user.Name, user.Email, user.Mobile, user.Password).Scan(&userValue).Error

	// return userValue, err

	createUserQuery := `INSERT INTO users(name,surname,email_id,password,phone,created_at)
				  VALUES($1,$2,$3,$4,$5,NOW())
				  RETURNING id,name,surname,email_id,phone`

	err := c.DB.Raw(createUserQuery, user.Name, user.Surname, user.EmailId, user.Password, user.Phone).Scan(&userData).Error

	if err == nil {
		//query for creating a new entry in the user_infos table.
		insertUserInfoQuery := `INSERT INTO user_infos (is_blocked,users_id) 
								VALUES('f',$1);`
		err = c.DB.Exec(insertUserInfoQuery, userData.ID).Error
	}

	// cartUseridQuery := `INSERT INTO carts(user_id)
	// 				   VALUES($1)`
	// err = c.DB.Exec(cartUseridQuery, userData.ID).Error

	return userData, err
}

func (c *userDatabase) FindByEmail(ctx context.Context, email string) (model.UserLoginVarifier, error) {
	var userData model.UserLoginVarifier
	findUserQuery := `SELECT u.id,u.name,u.surname,u.email_id,u.password,u.phone,info.is_blocked FROM users as u FULL OUTER JOIN user_infos as info ON u.id=info.users_id WHERE u.email_id=$1;`
	err := c.DB.Raw(findUserQuery, email).Scan(&userData).Error
	return userData, err

}

func (c *userDatabase) BlockUser(ctx context.Context, user_id string) (domain.UserInfo, error) {
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
func (c *userDatabase) UnBlockUser(ctx context.Context, user_id string) (domain.UserInfo, error) {
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

func (c *userDatabase) AddAddress(ctx context.Context, body model.AddressInput, userID string) (domain.Address, error) {
	var existingAddress, addAddress domain.Address

	findAddressQuery := `SELECT *FROM addresses WHERE users_id=$1`
	err := c.DB.Raw(findAddressQuery, userID).Scan(&existingAddress).Error
	if err != nil {
		return domain.Address{}, err

	}
	if existingAddress.ID == 0 || existingAddress.UsersID == "" {
		// no address is found in user table, insert query
		insertQuery := `INSERT INTO addresses(users_id,house_name,street,district,state,landmark,pin_code)
	 			   	VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING*;`
		err := c.DB.Raw(insertQuery, userID, body.HouseName, body.Street, body.District, body.State, body.Landmark, body.PinCode).Scan(&addAddress).Error
		return addAddress, err
	} else {
		// address already there,update it
		UpdateQuery := `UPDATE addresses SET house_name=$1,street=$2,district=$3,state=$4,landmark=$5,pin_code=$6 WHERE users_id=$7 RETURNING *;`
		err := c.DB.Raw(UpdateQuery, body.HouseName, body.Street, body.District, body.State, body.Landmark, body.PinCode, userID).Scan(&addAddress).Error
		return addAddress, err
	}
}

func (cr *userDatabase) ViewAddress(ctx context.Context, userID string) (domain.Address, error) {
	var address domain.Address
	viewAddressQuery := `SELECT *FROM addresses WHERE users_id=$1`
	err := cr.DB.Raw(viewAddressQuery, userID).Scan(&address).Error
	return address, err
}
