package handlers

import (
	"go-users/database"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteUser(db database.UserInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Fetching required variables from request
		params := mux.Vars(r)
		id := params["id"]

		// MongoDB call
		res, err := db.Delete(id)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Sending response
		WriteResponse(w, http.StatusOK, res)
	}
}
