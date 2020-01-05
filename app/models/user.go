package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Role string

const (
	// Has access to all features
	// Unsure about what more role pro will have
	Member Role = "member"
	Pro    Role = "pro"
)

type User struct {
	Model           // gorm.Model
	Email    string `gorm:"unique;not null" json:"email"`
	GoogleID string `gorm:"unique;not null" json:"google_id"`
	Role     Role   `gorm:"type:role;default:'member';not null;"`
}

// Used for oauth, either logs in user or signs them up
func (u *User) LoginOrSignup(db *gorm.DB) error {
	err := db.Model(User{}).Where("google_id = ?", u.GoogleID).Take(&u).Error

	// If the error is that the record is not found
	// sign them up and proceed
	if gorm.IsRecordNotFoundError(err) {
		if err := db.Create(&u).Error; err != nil {
			return err
		}
	} else if err != nil {
		// Handle other cases other then record not found
		return err
	}

	return nil
}

// Finds a single user, returns error
func (u *User) FindUserByID(db *gorm.DB, uid uint) error {
	err := db.Model(User{}).Where("id = ?", uid).Take(&u).Error

	// We want to identify if the user is not found for debugging purposes
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("User Not Found")
	} else if err != nil {
		return err
	}

	return nil
}
