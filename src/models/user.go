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
	Confirmed bool      `gorm:"not null;default:false"`
	CreatedAt time.Time `gorm:"not null"`
}

type Logging struct {
	ID              int       `gorm:"primaryKey"`
	Username        string    `gorm:"not null"`
	ConnectAt       time.Time `gorm:"not null"`
	FirstConnection bool      `gorm:"not null;default:false"`
}

type Token struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	TokenString string `json:"token"`
}
