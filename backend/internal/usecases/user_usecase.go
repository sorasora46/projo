package usecases

import (
	"github.com/google/uuid"
	"github.com/sorasora46/projo/backend/internal/adaptors/interfaces"
	"github.com/sorasora46/projo/backend/internal/dtos"
	"github.com/sorasora46/projo/backend/internal/entities"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	CreateUser(req dtos.CreateUserReq) error
	GetByUsername(username string) (*dtos.UserDTO, error)
	DeleteByUsername(username string) error
}

type UserService struct {
	repo interfaces.UserRepository
}

func NewUserUsercase(repo interfaces.UserRepository) UserUsecase {
	return &UserService{repo: repo}
}

func (u *UserService) CreateUser(req dtos.CreateUserReq) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser := entities.User{
		Id:             uuid.NewString(),
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}
	if err := u.repo.Create(newUser); err != nil {
		return err
	}
	return nil
}

func (u *UserService) GetByUsername(username string) (*dtos.UserDTO, error) {
	user, err := u.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	dto := dtos.UserDTO{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
	}

	return &dto, nil
}

func (u *UserService) DeleteByUsername(username string) error {
	if err := u.repo.DeleteByUsername(username); err != nil {
		return err
	}
	return nil
}
