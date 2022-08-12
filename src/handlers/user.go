package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MydroX/technical-test-fleurs-d-ici/pkg/headers"
	"github.com/MydroX/technical-test-fleurs-d-ici/pkg/response"
	"github.com/MydroX/technical-test-fleurs-d-ici/src/controllers"
	"github.com/MydroX/technical-test-fleurs-d-ici/src/models"
)

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

	headers.Set(w, http.StatusCreated)
	json.NewEncoder(w).Encode(response.Simple{Message: "account successfully created"})
}
