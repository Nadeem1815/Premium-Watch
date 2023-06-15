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

func (cr *adminDatabase) DashBoard(ctx context.Context) (model.AdminDashBoard, error) {
	var dashBoard model.AdminDashBoard
	dashBoardQuwery := `SELECT 
   							COUNT(CASE WHEN order_status_id = 1 THEN id END) AS completed_orders,
   							COUNT(CASE WHEN order_status_id = 3 THEN id END) AS pending_orders,
   							COUNT(CASE WHEN order_status_id = 2 THEN id END) AS cancelled_orders,
   							COUNT(id) AS total_orders,
   							SUM(CASE WHEN o.order_status_id != 2 AND o.order_status_id != 3 THEN o.order_total ELSE 0 END) AS order_value,
   							COUNT(DISTINCT o.user_id) AS ordered_users
						FROM 
   							 orders o;`
	err := cr.DB.Raw(dashBoardQuwery).Scan(&dashBoard).Error
	if err != nil {
		return dashBoard, err

	}
	totalOrderedItemsQuery := `SELECT COUNT(id)AS total_order_items FROM order_items`
	err = cr.DB.Raw(totalOrderedItemsQuery).Scan(&dashBoard.TotalOrderItems).Error
	if err != nil {
		return dashBoard, err

	}
	dashBoard.PendingAmount = dashBoard.OrderValue - dashBoard.CreditedAmount

	getTotalUsers := `SELECT 	COUNT(id)AS total_users FROM users`

	err = cr.DB.Raw(getTotalUsers).Scan(&dashBoard).Error
	if err != nil {
		return dashBoard, err

	}
	return dashBoard, nil
}

func (cr *adminDatabase) SalesRepo(ctx context.Context) ([]model.SalesReport, error) {
	var salesData []model.SalesReport
	saleDataReport := `SELECT 
									o.id AS order_id,
									o.user_id,
									o.order_total AS total,
									c.code AS coupon_code,
									pm.payment_method,
									os.order_status,
									ds.status AS delivery_status,
									o.order_date
						FROM 		
									orders o
						LEFT JOIN	
									payment_methods pm ON o.payment_method_id=pm.id
						LEFT JOIN	
									order_statuses os ON o.order_status_id=os.id
						LEFT JOIN
									delivery_statuses ds ON o.delivery_status_id=ds.id
						LEFT JOIN 
									coupons c ON o.coupon_id=c.id;`

	if err := cr.DB.Raw(saleDataReport).Scan(&salesData).Error; err != nil {
		return []model.SalesReport{}, err
	}
	// for _, sale := range salesData {
	// 	fmt.Println("Delivery Status:", sale.DeliveryStatus)
	// }
	return salesData, nil

}
