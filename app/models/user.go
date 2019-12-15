package models

type User struct {
	Model        // gorm.Model
	Email string `gorm:"unique;not null" json:"email"`
}
