package handlers

import (
	"encoding/json"
	"go-users/database"
	"go-users/models"
	"io/ioutil"
	"net/http"
	"time"
)

func InsertUser(db database.UserInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dummyUser := models.User{}

		// Reading Request Body in byte format
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		// Parsing request body to Golang objects
		err = json.Unmarshal(body, &dummyUser)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		// Creating new user model for inserting to mongoDB
		var user = models.User{
			ID:          dummyUser.ID,
			Name:        dummyUser.Name,
			DOB:         dummyUser.DOB,
			Address:     dummyUser.Address,
			Description: dummyUser.Description,
			CreatedAt:   time.Now(),
		}

		// MongoDB call
		res, err := db.Insert(user)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Sending response
		WriteResponse(w, http.StatusOK, res)
	}
}

// Wrapper function to send response
func WriteResponse(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}
