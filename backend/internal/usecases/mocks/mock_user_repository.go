package mocks

import (
	"github.com/sorasora46/projo/backend/internal/entities"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByUsername(username string) (*entities.User, error) {
	args := m.Called(username)
	if user := args.Get(0); user != nil {
		return user.(*entities.User), nil
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) DeleteByUsername(username string) error {
	args := m.Called(username)
	return args.Error(0)
}

func (m *MockUserRepository) GetLoginInfoByUsername(username string) (*entities.User, error) {
	args := m.Called(username)
	if user := args.Get(0); user != nil {
		return user.(*entities.User), nil
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) CheckIfUserExistByUniqueKey(uniqueKey string) (bool, error) {
	args := m.Called(uniqueKey)
	return args.Bool(0), args.Error(1)
}
