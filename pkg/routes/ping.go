package routes

import (
	controller "pkg/controller"

	"github.com/gorilla/mux"
)

func (m *Router) PingRoutes() *mux.Router {
	m.R.HandleFunc("/healthz", controller.HealthHandlerGET).Methods("GET")
	m.R.HandleFunc("/ping/{name}", controller.PingHandlerGET).Methods("GET")
	m.R.HandleFunc("/ping", controller.PingHandlerPOST).Methods("POST")
	m.R.HandleFunc("/ping/echo", controller.EchoHandlerPOST).Methods("POST")
	return m.R
}
