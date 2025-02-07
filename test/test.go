package tests

import (
	"GOUSERAPI/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) CreateUser(user models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockRepo.On("CreateUser", mock.Anything).Return(nil)
	assert.Nil(t, mockRepo.CreateUser(models.User{Name: "John", Email: "john@example.com"}))
}
