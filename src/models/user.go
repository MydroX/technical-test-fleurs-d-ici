package models

import "time"

type LoginBodyRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterBodyRequest struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type User struct {
	ID        int       `gorm:"primaryKey"`
	Username  string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}
