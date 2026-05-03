package entity

import (
	"time"
)

type User struct {
	ID          int        `json:"id" gorm:"column:id"`
	Name        string     `json:"name" gorm:"column:name"`
	Email       string     `json:"email" gorm:"column:email"`
	IsActive    bool       `json:"is_active" gorm:"column:is_active"`
	CreatedDate time.Time  `json:"created_date" gorm:"column:created_date"`
	UpdatedDate *time.Time `json:"updated_date" gorm:"column:updated_date"`
}
