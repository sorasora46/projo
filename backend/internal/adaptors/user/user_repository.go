package user

import "gorm.io/gorm"

type UserRepository interface {
	Create()
	Get()
	Update()
	Delete()
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db any) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (u *UserRepositoryImpl) Create() {}
func (u *UserRepositoryImpl) Get()    {}
func (u *UserRepositoryImpl) Update() {}
func (u *UserRepositoryImpl) Delete() {}
