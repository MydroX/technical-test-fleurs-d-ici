package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/MydroX/technical-test-fleurs-d-ici/pkg/db/postgres"
	"github.com/MydroX/technical-test-fleurs-d-ici/pkg/error"
	"github.com/MydroX/technical-test-fleurs-d-ici/pkg/password"
	"github.com/MydroX/technical-test-fleurs-d-ici/pkg/token"
	"github.com/MydroX/technical-test-fleurs-d-ici/src/models"
	"gorm.io/gorm"
)

func Login(request models.LoginBodyRequest) (*string, *error.Error) {
	if request.Username == "" || request.Password == "" {
		return nil, error.New(http.StatusBadRequest, "username and passwords are required")
	}

	var userDB models.User
	if err := postgres.Conn.Where("username = ?", request.Username).First(&userDB).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.New(http.StatusUnauthorized, "username or password is incorrect")
		}
		return nil, error.NewInternal("user login", fmt.Sprintf("error querying database: %v", err))
	}

	if !password.CheckPasswordHash(request.Password, userDB.Password) {
		return nil, error.New(http.StatusUnauthorized, "username or password is incorrect")
	} else if password.CheckPasswordHash(request.Password, userDB.Password) && request.Username == userDB.Username {

		var firstConnection bool

		err := postgres.Conn.Where("username = ?", request.Username).First(&models.Logging{Username: request.Username}).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				firstConnection = true
			} else {
				return nil, error.NewInternal("user login", fmt.Sprintf("error querying logs: %v", err))
			}
		}

		err = postgres.Conn.Create(&models.Logging{
			Username:        request.Username,
			ConnectAt:       time.Now(),
			FirstConnection: firstConnection,
		}).Error
		if err != nil {
			return nil, error.NewInternal("user login", fmt.Sprintf("error creating new log: %v", err))
		}

		token, err := token.GenerateJWT(userDB.Username)
		if err != nil {
			return nil, error.NewInternal("user login", fmt.Sprintf("error generating token: %v", err))
		}
		return &token, nil
	}

	return nil, error.NewInternal("user login", "unknown error")
}

func Register(request models.RegisterBodyRequest) *error.Error {
	if request.Username == "" || request.Password == "" || request.PasswordConfirmation == "" {
		return error.New(http.StatusBadRequest, "username and passwords are required")
	}

	if request.Password != request.PasswordConfirmation {
		return error.New(http.StatusBadRequest, "passwords must be the same")
	}

	var user models.User
	if err := postgres.Conn.Where("username = ?", request.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hashedPassword, err := password.HashPassword(request.Password)
			if err != nil {
				return error.NewInternal("user register", fmt.Sprintf("error hashing password: %v", err))
			}

			err = postgres.Conn.Create(&models.User{
				Username: request.Username,
				Password: hashedPassword,
			}).Error

			if err != nil {
				return error.NewInternal("user register", fmt.Sprintf("error creating new user: %v", err))
			}

			return nil
		}
		return error.NewInternal("user register", fmt.Sprintf("error querying database: %v", err))
	}

	if user.Username == request.Username && user.Password != request.Password {
		return error.New(http.StatusConflict, "username already exists")
	} else if user.Username == request.Username && user.Password == request.Password {
		return error.New(http.StatusConflict, "your account already exists")
	}

	return nil
}
