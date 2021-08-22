package main

import (
	"context"
	"go-users/config"
	"go-users/database"
	"go-users/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	conf := config.GetConfig()
	ctx := context.TODO()

	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)

	client := &database.UserClient{
		Col: collection,
		Ctx: ctx,
	}

	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.SearchUsers(client)).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUser(client)).Methods("GET")
	r.HandleFunc("/users", handlers.InsertUser(client)).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.UpdateUser(client)).Methods("PATCH")
	r.HandleFunc("/users/{id}", handlers.DeleteUser(client)).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
