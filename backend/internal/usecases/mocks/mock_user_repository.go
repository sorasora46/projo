package mocks

import "github.com/sorasora46/projo/backend/internal/entities"

type MockUserRepository struct {
	CreateFunc                      func(user entities.User) error
	GetByUsernameFunc               func(username string) (*entities.User, error)
	DeleteByUsernameFunc            func(username string) error
	GetHashedPasswordByUsernameFunc func(username string) ([]byte, error)
}

func (m *MockUserRepository) Create(user entities.User) error {
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

func (m *MockUserRepository) GetHashedPasswordByUsername(username string) ([]byte, error) {
	if m.GetHashedPasswordByUsernameFunc != nil {
		return m.GetHashedPasswordByUsernameFunc(username)
	}
	return nil, nil
}
