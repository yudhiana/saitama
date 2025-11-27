package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Email     string         `gorm:"type:varchar(255);unique;not null" json:"email"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (User) TableName() string {
	return "users"
}
