package repository

import (
	"context"
	"fmt"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

//-----------------------------UserSignUp----------------------

func (c *userDatabase) UserSignUp(ctx context.Context, user req.UserReq) (res.UserData, error) {
	var userData res.UserData
	insertQuery := `INSERT INTO users (name,email,mobile,password)VALUES($1,$2,$3,$4) 
					RETURNING id,name,email,mobile`
	err := c.DB.Raw(insertQuery, user.Name, user.Email, user.Mobile, user.Password).Scan(&userData).Error
	return userData, err
}

//-----------------------------UserLogin------------------------

func (c *userDatabase) UserLogin(ctx context.Context, email string) (domain.Users, error) {
	var userData domain.Users
	err := c.DB.Raw("SELECT * FROM users WHERE email=?", email).Scan(&userData).Error
	return userData, err
}
func (c *userDatabase) IsSignIn(phone string) (bool, error) {
	query := "select exists(select 1 from users where mobie=?)"
	var IsSignIn bool
	err := c.DB.Raw(query, phone).Scan(&IsSignIn).Error
	return IsSignIn, err
}

func (c *userDatabase) OtpLogin(phone string) (int, error) {
	var id int
	query := "SELECT id from users where mobile=?"
	err := c.DB.Raw(query, phone).Scan(&id).Error
	return id, err
}
func (c *userDatabase) AddAddress(address req.Address) error {
	query:=`INSERT INTO address FROM  `
	c.DB.Raw(query,)
	return fmt.Errorf("")
}
