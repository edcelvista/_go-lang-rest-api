package routes

import (
	controller "pkg/controller"

	"github.com/gorilla/mux"
)

func (m *Router) CrudRoutes() *mux.Router {
	m.R.HandleFunc("/crud/list", controller.CrudHandlerLIST).Methods("GET")
	m.R.HandleFunc("/crud/{messageId}", controller.CrudHandlerGET).Methods("GET")
	m.R.HandleFunc("/crud", controller.CrudHandlerPOST).Methods("POST")
	m.R.HandleFunc("/crud/{messageId}", controller.CrudHandlerDELETE).Methods("DELETE")
	return m.R
}
