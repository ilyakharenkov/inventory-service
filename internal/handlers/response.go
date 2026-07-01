package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func BaseResponse(data interface{}, w http.ResponseWriter, status int) {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}
