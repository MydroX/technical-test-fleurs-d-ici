package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Payload interface{}
}

type Simple struct {
	Message string `json:"message"`
}

func (r *Response) JSON(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(r)
}
