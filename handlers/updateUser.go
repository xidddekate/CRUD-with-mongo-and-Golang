package handlers

import (
	"encoding/json"
	"go-users/database"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateUser(db database.UserInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Fetching required variables from request
		params := mux.Vars(r)
		id := params["id"]

		// Reading Request Body in byte format
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Parsing request body to Golang objects
		var user interface{}
		err = json.Unmarshal(body, &user)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// MongoDB call
		res, err := db.Update(id, user)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// sending response
		WriteResponse(w, http.StatusOK, res)
	}
}
