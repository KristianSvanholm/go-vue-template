package customhttp

import (
	"encoding/json"
	"net/http"
)

// Encode any struct given to it
func Encode(w http.ResponseWriter, data interface{}) error {
	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}
