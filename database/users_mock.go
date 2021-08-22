package database

import (
	"fmt"
	"go-users/models"

	"github.com/stretchr/testify/mock"
)

type MockUserClient struct {
	mock.Mock
}

// Mock Insert Function to mock actual function
func (m *MockUserClient) Insert(user models.User) (models.User, error) {
	// Setting variable arguments
	args := m.Called(user)
	// returning User model and error
	return args.Get(0).(models.User), args.Error(1)
}

// Mock Update Function to mock actual function
func (m *MockUserClient) Update(id string, update interface{}) (models.UserUpdate, error) {
	// Setting variable arguments
	args := m.Called(id, update)
	// returning UserUpdate model and error
	return args.Get(0).(models.UserUpdate), args.Error(1)
}

// Mock Delete Function to mock actual function
func (m *MockUserClient) Delete(id string) (models.UserDelete, error) {
	// Setting variable arguments
	args := m.Called(id)
	// returning UserDelete model and error
	return args.Get(0).(models.UserDelete), args.Error(1)
}

// Mock Get Function to mock actual function
func (m *MockUserClient) Get(id string) (models.User, error) {
	fmt.Println("call get mock function")
	// Setting variable arguments
	args := m.Called(id)
	// returning User model and error
	return args.Get(0).(models.User), args.Error(1)
}

// Mock Search Function to mock actual function
func (m *MockUserClient) Search(filter interface{}) ([]models.User, error) {
	// Setting variable arguments
	args := m.Called(filter)
	// returning User model and error
	return args.Get(0).([]models.User), args.Error(1)
}
