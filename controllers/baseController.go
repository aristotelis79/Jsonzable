package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterBaseRoutes adds base routes to router
func RegisterBaseRoutes(r *mux.Router, p string) {
	ur := r.PathPrefix(p).Subrouter()
	// swagger:route GET /version
	// ---
	// Gets version.
	// responses:
	//   200: version
	ur.HandleFunc("/version", GetVersion).Methods("Get")
}

type baseResponse struct {
	Code    int    `json:"code"`
	Data    string `json:"data,omitempty"`
	Message string `json:"msg,omitempty"`
}

// GetVersion ...
func GetVersion(w http.ResponseWriter, r *http.Request) {
	Response(w, "v1")
}

// Response ...
//func Response(w http.ResponseWriter, data map[string]interface{}) {
func Response(w http.ResponseWriter, data string) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
