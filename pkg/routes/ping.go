package routes

import (
	Controller "pkg/controller"

	"github.com/gorilla/mux"
)

func (m *Router) PingRoutes() *mux.Router {
	m.R.HandleFunc("/healthz", Controller.HealthHandlerGET).Methods("GET")
	m.R.HandleFunc("/ping/{name}", Controller.PingHandlerGET).Methods("GET")
	m.R.HandleFunc("/ping", Controller.PingHandlerPOST).Methods("POST")
	m.R.HandleFunc("/ping/echo", Controller.EchoHandlerPOST).Methods("POST")
	return m.R
}
