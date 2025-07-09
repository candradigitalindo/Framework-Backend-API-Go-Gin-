package models

import (
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Role struct {
	ID        string    `json:"id" gorm:"primaryKey;type:char(26)"`
	Name      string    `json:"name" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = ulid.Make().String()
	return
}
