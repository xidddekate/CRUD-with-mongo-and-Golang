package handlers_test

import (
	"go-users/database"
	"go-users/handlers"
	"go-users/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

func TestUpdateUser(t *testing.T) {
	// Instance of mock user client
	client := &database.MockUserClient{}
	id := primitive.NewObjectID().Hex()

	// Testcases
	tests := map[string]struct {
		id           string
		payload      string
		expectedCode int
	}{
		"should return 200": {
			id:           id,
			payload:      `{"name":"SUperMan","description": "Manavv Test"}`,
			expectedCode: 200,
		},
		"should return 404": {
			id:           "",
			expectedCode: 404,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// mock the update user function
			if test.expectedCode == 200 {
				client.On("Update", test.id, mock.Anything).Return(models.UserUpdate{}, nil)
			}
			req, _ := http.NewRequest("PATCH", "/users/"+test.id, strings.NewReader(test.payload))
			rec := httptest.NewRecorder()

			// make request to update API and record response
			r := mux.NewRouter()
			r.HandleFunc("/users/{id}", handlers.UpdateUser(client))
			r.ServeHTTP(rec, req)

			// Assert Expectations from testify mock package
			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Update")
			}
		})
	}
}
