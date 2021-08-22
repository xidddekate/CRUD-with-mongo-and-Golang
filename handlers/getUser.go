package handlers

import (
	"go-users/database"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUser(db database.UserInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Fetching required variables from request
		params := mux.Vars(r)
		id := params["id"]

		// MongoDB call
		res, err := db.Get(id)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Sending response
		WriteResponse(w, http.StatusOK, res)
	}
}
