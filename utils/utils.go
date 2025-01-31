package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error{
	if r.Body == nil{
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error{
	w.Header().Add("content_typo", "aplication/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriterError(w http.ResponseWriter, status int,  err error){
	WriteJSON(w, status, map[string]string{"error":err.Error()})
}