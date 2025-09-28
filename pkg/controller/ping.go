package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	model "pkg/model"
)

func HealthHandlerGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := &model.HealthRes{Status: "Running"}
	json.NewEncoder(w).Encode(res)
}

func PingHandlerGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqParams := mux.Vars(r)

	res := &model.Pong{Headers: map[string]string{"Headers": fmt.Sprintf("%+v", r.Header), "Host": r.Host, "Source": r.RemoteAddr}, Pong: "pong", Message: reqParams}
	json.NewEncoder(w).Encode(res)
}

func PingHandlerPOST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var parsedUserRequest model.Ping
	err := json.NewDecoder(r.Body).Decode(&parsedUserRequest)
	if err != nil {
		log.Printf("‼️ Invalid JSON %v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	res := &model.Pong{Headers: map[string]string{"Headers": fmt.Sprintf("%+v", r.Header), "Host": r.Host, "Source": r.RemoteAddr}, Pong: "pong", Message: parsedUserRequest.Message}
	json.NewEncoder(w).Encode(res)
}

func EchoHandlerPOST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("‼️ Invalid JSON %v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	res := &model.Echo{EchoHeaders: map[string]string{"Headers": fmt.Sprintf("%+v", r.Header), "Host": r.Host, "Source": r.RemoteAddr}, EchoData: data}
	json.NewEncoder(w).Encode(res)
}
