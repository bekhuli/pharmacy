package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func BindJSON(r *http.Request, dst any) error {
	defer r.Body.Close()

	if r.Header.Get("Content-Type") != "application/json" {
		return errors.New("invalid content type")
	}

	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		return fmt.Errorf("invalid json body: %w", err)
	}

	return nil
}
