package handlers_test

import (
	"go-users/database"
	"go-users/handlers"
	"go-users/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/gorilla/mux"
)

func TestSearchUsers(t *testing.T) {
	// Instance of mock user client
	client := &database.MockUserClient{}

	// Testcases
	tests := map[string]struct {
		payload      string
		expectedCode int
	}{
		"should return 200": {
			payload:      `{"description":"Man Test"}`,
			expectedCode: 200,
		},
		"should return 400": {
			payload:      "invalid json string",
			expectedCode: 400,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {

			// mock the search user by filter function
			client.On("Search", mock.Anything).Return([]models.User{}, nil)

			req, _ := http.NewRequest("GET", "/users?q="+test.payload, nil)
			rec := httptest.NewRecorder()

			// make request to Search Users by filter API and record response
			r := mux.NewRouter()
			r.HandleFunc("/users", handlers.SearchUsers(client))
			r.ServeHTTP(rec, req)

			// Assert Expectations from testify mock package
			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Search")
			}

		})
	}

}
