package routes

import (
	Controller "pkg/controller"

	"github.com/gorilla/mux"
)

func (m *Router) CrudRoutes() *mux.Router {
	m.R.HandleFunc("/crud/list", Controller.CrudHandlerLIST).Methods("GET")
	m.R.HandleFunc("/crud/{messageId}", Controller.CrudHandlerGET).Methods("GET")
	m.R.HandleFunc("/crud", Controller.CrudHandlerPOST).Methods("POST")
	m.R.HandleFunc("/crud/{messageId}", Controller.CrudHandlerDELETE).Methods("DELETE")
	return m.R
}
