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

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err = json.Unmarshal(body, &dummyUser)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var user = models.User{
			ID:          dummyUser.ID,
			Name:        dummyUser.Name,
			DOB:         dummyUser.DOB,
			Address:     dummyUser.Address,
			Description: dummyUser.Description,
			CreatedAt:   time.Now(),
		}

		res, err := db.Insert(user)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(w, http.StatusOK, res)
	}
}

func WriteResponse(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}
