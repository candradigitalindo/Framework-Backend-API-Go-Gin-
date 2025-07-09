package models

import (
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;type:char(26)"`
	Name      string    `json:"name" gorm:"not null"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	RoleID    string    `json:"role_id"`
	Role      Role      `json:"role" gorm:"foreignKey:RoleID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BeforeCreate hook to set ULID before inserting a new record
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = ulid.Make().String()
	return
}
