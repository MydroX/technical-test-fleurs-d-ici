package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/MydroX/technical-test-fleurs-d-ici/pkg/db/postgres"
	"github.com/MydroX/technical-test-fleurs-d-ici/pkg/error"
	"github.com/MydroX/technical-test-fleurs-d-ici/src/models"
	"gorm.io/gorm"
)

func Register(request models.RegisterBodyRequest) *error.Error {
	if request.Username == "" || request.Password == "" {
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
