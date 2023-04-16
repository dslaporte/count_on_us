package pkg_webserver

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

func MalformedBody(w http.ResponseWriter, errors []string) {
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(&ResponseError{
		Message: "Malformed Body",
		Errors:  errors,
	})
}

func InternalServerError(w http.ResponseWriter, errors []string) {
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(&ResponseError{
		Message: "Internal Server Error",
		Errors:  errors,
	})
}

func Created(w http.ResponseWriter, response any) {
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(&response)
}

func OK(w http.ResponseWriter, response any) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&response)
}

func NotFound(w http.ResponseWriter, errors []string) {
	w.WriteHeader(http.StatusNotFound)
	_ = json.NewEncoder(w).Encode(&ResponseError{
		Message: "Record not found",
		Errors:  errors,
	})
}
