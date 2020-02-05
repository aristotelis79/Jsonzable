package routers

import (
	"github.com/aristotelis79/Jsonzable/controllers"
	"github.com/gorilla/mux"
)

// RegisterV1Routes ...
func RegisterV1Routes(r *mux.Router) {
	v1 := r.PathPrefix("/v1").Subrouter()
	controllers.RegisterBaseRoutes(v1, "/base")
}
