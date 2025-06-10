package user

import (
	"github.com/sorasora46/projo/backend/internal/adaptors/interfaces"
	"github.com/sorasora46/projo/backend/internal/entities"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (u *UserRepositoryImpl) Create(user *entities.User) error {
	transaction := u.db.Create(user)
	if transaction.Error != nil {
		return transaction.Error
	}
	return nil
}

func (u *UserRepositoryImpl) GetByUsername(username string) (*entities.User, error) {
	var user entities.User
	transaction := u.db.Where("username = ?", username).Omit("hashed_password").First(&user)
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	return &user, nil
}

func (u *UserRepositoryImpl) DeleteByUsername(username string) error {
	transaction := u.db.Where("username = ?", username).Delete(&entities.User{})
	if transaction.Error != nil {
		return transaction.Error
	}
	return nil
}

func (u *UserRepositoryImpl) GetLoginInfoByUsername(username string) (*entities.User, error) {
	var user entities.User
	transaction := u.db.Where("username = ?", username).First(&user)
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	return &user, nil
}

func (u *UserRepositoryImpl) CheckIfUserExistByUniqueKey(uniqueKey string) (bool, error) {
	var count int64
	u.db.Where("id = ?", uniqueKey).Or("username = ?", uniqueKey).Or("email = ?", uniqueKey).Find(&entities.User{}).Count(&count)
	return count == 1, nil
}
