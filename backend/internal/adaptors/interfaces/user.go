package interfaces

import "github.com/sorasora46/projo/backend/internal/entities"

type UserRepository interface {
	Create(user *entities.User) error
	GetByUsername(username string) (*entities.User, error)
	DeleteByUsername(username string) error
	GetLoginInfoByUsername(username string) (*entities.User, error)
	CheckIfUserExistByUniqueKey(uniqueKey string) (bool, error)
}
