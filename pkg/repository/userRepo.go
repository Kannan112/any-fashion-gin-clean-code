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

// UserSignUp creates a new user account.
func (c *userDatabase) UserSignUp(ctx context.Context, user req.UserReq) (res.UserData, error) {
	var userData res.UserData
	insertQuery := `
		INSERT INTO users (name, email, mobile, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, email, mobile
	`
	err := c.DB.Raw(insertQuery, user.Name, user.Email, user.Mobile, user.Password).Scan(&userData).Error
	return userData, err
}

// UserLogin retrieves a user's data based on the provided email.
func (c *userDatabase) UserLogin(ctx context.Context, email string) (domain.Users, error) {
	var userData domain.Users
	var userBlocked bool

	// Check if the user is blocked by the admin.
	query := `SELECT EXISTS(SELECT * FROM users WHERE email=$1 AND is_blocked=true)`
	err := c.DB.Raw(query, email).Scan(&userBlocked).Error
	if err != nil {
		return userData, err
	}
	if userBlocked {
		return userData, fmt.Errorf("user is blocked")
	}

	err = c.DB.Raw("SELECT * FROM users WHERE email=?", email).Scan(&userData).Error
	return userData, err
}

// IsSignIn checks if a user is signed in based on their phone number.
func (c *userDatabase) IsSignIn(phone string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE mobile=?)"
	var isSignIn bool
	err := c.DB.Raw(query, phone).Scan(&isSignIn).Error
	return isSignIn, err
}

// OtpLogin retrieves the user ID based on the provided phone number.
func (c *userDatabase) OtpLogin(phone string) (int, error) {
	var id int
	query := "SELECT id FROM users WHERE mobile=?"
	err := c.DB.Raw(query, phone).Scan(&id).Error
	return id, err
}

// AddAddress adds a new address for a user.
func (c *userDatabase) AddAddress(id int, address req.AddAddress) error {
	// Check if the address is set as the default address.
	if address.IsDefault {
		changeAddress := `
			UPDATE addresses
			SET is_default=$1
			WHERE users_id=$2 AND is_default=$3
		`
		err := c.DB.Exec(changeAddress, false, id, true).Error
		if err != nil {
			fmt.Println("SET1")
		}
	}

	query := `
		INSERT INTO addresses (users_id, house_number, street, city, district, landmark, pincode, is_default)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	err := c.DB.Exec(query, id, address.House_number, address.Street, address.City, address.District, address.Landmark, address.Pincode, address.IsDefault).Error
	return err
}

// UpdateAddress updates an existing user's address.
func (c *userDatabase) UpdateAddress(id int, addressID int, address req.AddAddress) error {
	// Check if the address should be set as the default address.
	if address.IsDefault {
		changeDefault := `
			UPDATE addresses
			SET is_default = $1
			WHERE users_id = $2 AND is_default = $3
		`
		err := c.DB.Exec(changeDefault, false, id, true).Error

		if err != nil {
			return err
		}
	}

	// Check if the address belongs to the user.
	var check bool
	addressExists := `
		SELECT EXISTS(SELECT * FROM addresses WHERE users_id = $1 AND id = $2)
	`
	err := c.DB.Raw(addressExists, id, addressID).Scan(&check).Error
	if err != nil {
		return err
	}
	if !check {
		return fmt.Errorf("wrong address id")
	}

	// Update the address.
	updateQuery := `
		UPDATE addresses
		SET users_id = $1, house_number = $2, street = $3, city = $4, district = $5, landmark = $6, pincode = $7, is_default = $8
		WHERE users_id = $9 AND id = $10
	`
	err = c.DB.Exec(updateQuery, id, address.House_number, address.Street, address.City, address.District, address.Landmark, address.Pincode, address.IsDefault, id, addressID).Error
	return err
}

// ViewProfile retrieves the user's profile based on the provided user ID.
func (c *userDatabase) ViewProfile(id int) (res.UserData, error) {
	var profile res.UserData
	findProfile := `
		SELECT name, email, mobile
		FROM users
		WHERE id = $1
	`
	err := c.DB.Raw(findProfile, id).Scan(&profile).Error
	fmt.Println(profile)
	return profile, err
}

// EditProfile updates the user's profile based on the provided user ID and updated details.
func (c *userDatabase) EditProfile(id int, updatingDetails req.UserReq) (res.UserData, error) {
	var profile res.UserData
	tx := c.DB.Begin()
	fmt.Println("update Test in Edit profile: ", updatingDetails.Name)
	updatedQuery := `
		UPDATE users
		SET name = $1, email = $2, mobile = $3
		WHERE id = $4
		RETURNING name, email, mobile
	`
	err := tx.Raw(updatedQuery, updatingDetails.Name, updatingDetails.Email, updatingDetails.Mobile, id).Scan(&profile).Error
	if err != nil {
		tx.Rollback()
		return profile, err
	}
	fmt.Println("EditProfile test 1")
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return profile, err
	}
	return profile, err
}

// ListAllAddress lists all the addresses associated with a user.
func (c *userDatabase) ListAllAddress(id int) ([]domain.Addresss, error) {
	var list []domain.Addresss
	query := `
		SELECT *
		FROM addresses
		WHERE users_id = $1
	`
	err := c.DB.Raw(query, id).Scan(&list).Error
	return list, err
}

// DeleteAddress deletes the specified address for a user.
func (c *userDatabase) DeleteAddress(ctx context.Context, userID, addressID int) ([]domain.Addresss, error) {
	var domain []domain.Addresss
	var check bool
	tx := c.DB.Begin()
	addressExists := `
		SELECT EXISTS(SELECT id FROM addresses WHERE id = $1)
	`
	err := tx.Raw(addressExists, addressID).Scan(&check).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if !check {
		tx.Rollback()
		return nil, fmt.Errorf("please enter a valid address id")
	}
	deleteQuery := `
		DELETE FROM addresses
		WHERE users_id = $1 AND id = $2
	`
	err = tx.Exec(deleteQuery, userID, addressID).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return domain, err
}

// FindAddress checks if the user has a default address.
func (c *userDatabase) FindAddress(ctx context.Context, userID int) (bool, error) {
	var exists bool
	checkAddress := `
		SELECT EXISTS(SELECT * FROM addresses WHERE users_id = $1 AND is_default = true)
	`
	err := c.DB.Raw(checkAddress, userID).Scan(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}
