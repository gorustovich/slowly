package app

import (
	"net/http"
	"encoding/json"
	"fmt"
	"log"
)

func NewSuccessResponse(message string) response {
	return response{
		Status: message,
		Code:   http.StatusOK,
	}
}

type response struct {
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
	Code   int    `json:"-"`
}

func NewErrorResponse(message string, code int) response {
	return response{
		Error: message,
		Code:  code,
	}
}

func (resp response) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)

	var b []byte
	if resp.Status != "" || resp.Error != "" {
		var err error
		b, err = json.Marshal(resp)
		if err != nil {
			log.Printf("can not marshal error response, error: %s\n", err)
			http.Error(w, "some troubles, try else or wait", http.StatusInternalServerError)
		}
		fmt.Fprintln(w, string(b))
	}
}
