package handlers

import (
	"encoding/json"
	"go-users/database"
	"net/http"
)

func SearchUsers(db database.UserInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var filter interface{}

		// getting query from request url
		query := r.URL.Query().Get("q")

		// Parsing query to Golang object
		if query != "" {
			err := json.Unmarshal([]byte(query), &filter)
			if err != nil {
				WriteResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		}

		// MongoDB call
		res, err := db.Search(filter)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// sending response
		WriteResponse(w, http.StatusOK, res)
	}
}
