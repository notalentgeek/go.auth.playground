package models

import (
	"gorm.io/gorm"
)

// Users ...
type Users struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Address  string `json:"address" db:"address"`
}

// UsersAdapter ...
type UsersAdapter struct {
	DB *gorm.DB
}

// Create ...
func (u *UsersAdapter) Create(users Users) error {
	u.DB.Create(&users)

	return nil
}

// GetByName ...
func (u *UsersAdapter) GetByName(name string) (*Users, error) {
	user := &Users{}
	u.DB.Where("username = ?", name).First(user)

	return user, nil
}
