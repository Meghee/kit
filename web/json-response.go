package web

import (
	"encoding/json"
	"net/http"
)

// JSONResponse writes data as json to the response writer.
func JSONResponse(w http.ResponseWriter, status int, body interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&body)
}

// JSONSuccessResponse writes a success json response message to the
// response writer.
func JSONSuccessResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	body := map[string]string{
		"status":  "success",
		"message": message,
	}
	json.NewEncoder(w).Encode(body)
}

// JSONErrorResponse writes an error json response message to the
// response writer.
func JSONErrorResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	body := map[string]string{
		"status":  "error",
		"message": message,
	}
	json.NewEncoder(w).Encode(body)
}

// JSONWarningResponse writes a warning json response message to the
// response writer.
func JSONWarningResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	body := map[string]string{
		"status":  "warning",
		"message": message,
	}
	json.NewEncoder(w).Encode(body)
}
