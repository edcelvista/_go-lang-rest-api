package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	lib "pkg/lib"
	model "pkg/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client = lib.InitDatabase()

func CrudHandlerLIST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := lib.DbProfile{
		Client:   client,
		Action:   "list",
		Database: "golangapi",
		Table:    "messages",
	}

	data, err := db.Execute()
	if err != nil {
		log.Printf("‼️ Failed Listing Records %v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := model.CrudResWithBody{Action: r.Method, Response: data}
	json.NewEncoder(w).Encode(res)
}

func CrudHandlerGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqParams := mux.Vars(r)

	objID, err := primitive.ObjectIDFromHex(reqParams["messageId"])
	if err != nil {
		log.Printf("‼️ Invalid ObjectID converstion from hex %v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := lib.DbProfile{
		Client:   client,
		Action:   "find",
		Database: "golangapi",
		Table:    "messages",
		Filter:   bson.M{"_id": objID},
	}

	data, err := db.Execute()
	if err != nil {
		log.Printf("‼️ Failed Fetching Record %v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := model.CrudResWithBody{Action: r.Method, MessageId: reqParams["messageId"], Response: data}
	json.NewEncoder(w).Encode(res)
}

func CrudHandlerPOST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var parsedRequestBody model.CrudPostReq
	parsedRequestBody.Dt = time.Now().Unix()
	err := json.NewDecoder(r.Body).Decode(&parsedRequestBody)
	if err != nil {
		log.Printf("‼️ Invalid JSON %v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	db := lib.DbProfile{
		Client:   client,
		Action:   "insert",
		Database: "golangapi",
		Table:    "messages",
		Payload:  parsedRequestBody,
	}

	objectId, err := db.Execute()
	if err != nil {
		log.Printf("‼️ Failed Inserting Record %v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := model.CrudRes{Action: r.Method, MessageId: objectId.(string)}
	json.NewEncoder(w).Encode(res)
}

func CrudHandlerDELETE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqParams := mux.Vars(r)

	objID, err := primitive.ObjectIDFromHex(reqParams["messageId"])
	if err != nil {
		log.Printf("‼️ Invalid MessageId %v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := lib.DbProfile{
		Client:   client,
		Action:   "delete",
		Database: "golangapi",
		Table:    "messages",
		Filter:   bson.M{"_id": objID},
	}

	data, err := db.Execute()
	if err != nil {
		log.Printf("‼️ Failed Delete Record %v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := model.CrudResWithBody{Action: r.Method, MessageId: reqParams["messageId"], Response: data}
	json.NewEncoder(w).Encode(res)
}
