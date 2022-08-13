package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MydroX/technical-test-fleurs-d-ici/pkg/headers"
	"github.com/MydroX/technical-test-fleurs-d-ici/src/controllers"
	"github.com/MydroX/technical-test-fleurs-d-ici/src/models"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var body models.LoginBodyRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding body: %v", err), http.StatusBadRequest)
		return
	}

	token, er := controllers.Login(body)
	if er != nil {
		er.JSON(w)
		return
	}

	headers.Set(w, http.StatusOK)
	json.NewEncoder(w).Encode(LoginResponse{Token: *token})
}

type RegisterResponse struct {
	Message string `json:"message"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var body models.RegisterBodyRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding body: %v", err), http.StatusBadRequest)
		return
	}

	res := controllers.Register(body)
	if res != nil {
		res.JSON(w)
		return
	}

	//TODO: send by email his confirmation token to confirm his account

	headers.Set(w, http.StatusCreated)
	json.NewEncoder(w).Encode(RegisterResponse{Message: "user created"})
}
