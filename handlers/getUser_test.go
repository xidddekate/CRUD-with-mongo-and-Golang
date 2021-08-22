package handlers_test

import (
	"go-users/database"
	"go-users/handlers"
	"go-users/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"
)

func TestGetUser(t *testing.T) {
	// Instance of mock user client
	client := &database.MockUserClient{}
	id := primitive.NewObjectID().Hex()

	// testcases
	tests := map[string]struct {
		id           string
		expectedCode int
	}{
		"should return 200": {
			id:           id,
			expectedCode: 200,
		},
		"should return 404": {
			id:           "",
			expectedCode: 404,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// mock the get function
			if test.expectedCode == 200 {
				client.On("Get", test.id).Return(models.User{}, nil)
			}
			req, _ := http.NewRequest("GET", "/users/"+test.id, nil)
			rec := httptest.NewRecorder()

			// make request to get Users API and record response
			r := mux.NewRouter()
			r.HandleFunc("/users/{id}", handlers.GetUser(client))
			r.ServeHTTP(rec, req)

			// Assert Expectations from testify mock package
			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Get")
			}
		})
	}
}
