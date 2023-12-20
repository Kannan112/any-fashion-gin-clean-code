package domain

import "time"

type Users struct {
	ID        uint      `gorm:"primaryKey;unique;not null"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required,email" gorm:"unique;not null"`
	Mobile    string    `json:"mobile" binding:"required,eq=10" gorm:"unique"`
	Images    string    `json:"images"`
	Verify    bool      `json:"verify" gorm:"default:false"`
	Password  string    `json:"password"`
	IsBlocked bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type UsersData struct {
	ID        uint   `gorm:"primaryKey;unique;not null"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email" gorm:"unique;not null"`
	Mobile    string `json:"mobile" binding:"required,eq=10" gorm:"unique; not null"`
	IsBlocked bool   `gorm:"default:false"`
	CreatedAt time.Time
}

type Address struct {
	ID           uint `gorm:"primaryKey;unique;not null"`
	UsersID      uint
	Users        Users  `gorm:"foreignKey:UsersID"`
	House_number string `json:"house_number" binding:"required"`
	Street       string `json:"street" binding:"required"`
	City         string `json:"city" binding:"required"`
	District     string `json:"district " binding:"required"`
	Landmark     string `json:"landmark" binding:"required"`
	Pincode      int    `json:"pincode " binding:"required"`
	IsDefault    bool   `gorm:"default:false"`
}
type Addresss struct {
	ID           uint   `gorm:"primaryKey;unique;not null"`
	House_number string `json:"house_number" binding:"required"`
	Street       string `json:"street" binding:"required"`
	City         string `json:"city" binding:"required"`
	District     string `json:"district " binding:"required"`
	Landmark     string `json:"landmark" binding:"required"`
	Pincode      int    `json:"pincode " binding:"required"`
	IsDefault    bool   `gorm:"default:false"`
}

type UserInfo struct {
	ID                uint `gorm:"primaryKey"`
	UsersID           uint
	Users             Users `gorm:"foreignKey:UsersID"`
	BlockedAt         time.Time
	BlockedBy         uint
	ReasonForBlocking string
}
