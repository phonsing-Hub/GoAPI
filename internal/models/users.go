package models

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

