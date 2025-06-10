package mocks

import (
	"github.com/sorasora46/projo/backend/internal/entities"
)

type MockUserRepository struct {
	CreateFunc                      func(user *entities.User) error
	GetByUsernameFunc               func(username string) (*entities.User, error)
	DeleteByUsernameFunc            func(username string) error
	GetLoginInfoByUsernameFunc      func(username string) (*entities.User, error)
	CheckIfUserExistByUniqueKeyFunc func(uniqueKey string) (bool, error)
}

func (m *MockUserRepository) Create(user *entities.User) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(user)
	}
	return nil
}

func (m *MockUserRepository) GetByUsername(username string) (*entities.User, error) {
	if m.GetByUsernameFunc != nil {
		return m.GetByUsernameFunc(username)
	}
	return nil, nil
}

func (m *MockUserRepository) DeleteByUsername(username string) error {
	if m.DeleteByUsernameFunc != nil {
		return m.DeleteByUsernameFunc(username)
	}
	return nil
}

func (m *MockUserRepository) GetLoginInfoByUsername(username string) (*entities.User, error) {
	if m.GetLoginInfoByUsernameFunc != nil {
		return m.GetLoginInfoByUsernameFunc(username)
	}
	return nil, nil
}

func (m *MockUserRepository) CheckIfUserExistByUniqueKey(uniqueKey string) (bool, error) {
	if m.CheckIfUserExistByUniqueKeyFunc != nil {
		return m.CheckIfUserExistByUniqueKeyFunc(uniqueKey)
	}
	return false, nil
}
