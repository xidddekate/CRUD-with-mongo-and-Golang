package handlers_test

import (
	"go-users/database"
	"go-users/handlers"
	"go-users/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestInsertUser(t *testing.T) {
	// Instance of mock user client
	client := &database.MockUserClient{}
	// testcases
	tests := map[string]struct {
		payload      string
		expectedCode int
	}{
		"should return 200": {
			payload: `{
				"name":"Man",
				"dob":"2021-08-21",
				"description": "Man Test",
				"address" : "something"
			}`,
			expectedCode: 200,
		},
		"should return 400": {
			payload:      "invalid json string",
			expectedCode: 400,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// mock the insert function
			client.On("Insert", mock.Anything).Return(models.User{}, nil)
			req, _ := http.NewRequest("POST", "/users", strings.NewReader(test.payload))
			rec := httptest.NewRecorder()

			// make request to insert API and record response
			h := http.HandlerFunc(handlers.InsertUser(client))
			h.ServeHTTP(rec, req)

			// Assert Expectations from testify mock package
			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Insert")
			}
		})
	}
}
