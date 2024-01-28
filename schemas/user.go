package schemas

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Openings []Opening `gorm:"foreignkey:UserID"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Openings  []Opening `json:"openings"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
