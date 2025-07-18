package usecases

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sorasora46/projo/backend/config"
	"github.com/sorasora46/projo/backend/internal/adaptors/interfaces"
	"github.com/sorasora46/projo/backend/internal/dtos"
	"github.com/sorasora46/projo/backend/internal/dtos/req"
	"github.com/sorasora46/projo/backend/internal/entities"
	"github.com/sorasora46/projo/backend/pkg/constants"
	"github.com/sorasora46/projo/backend/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	CreateUser(req req.CreateUserReq) error
	GetByUsername(username string) (*dtos.UserDTO, error)
	DeleteByUsername(username string) error
	Login(username string, password string) (*string, error)
}

type UserService struct {
	envManager config.EnvManager
	repo       interfaces.UserRepository
}

func NewUserUsercase(repo interfaces.UserRepository, envManager config.EnvManager) UserUsecase {
	return &UserService{repo: repo, envManager: envManager}
}

func (u *UserService) CreateUser(req req.CreateUserReq) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser := &entities.User{
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

func (u *UserService) Login(username string, password string) (*string, error) {
	user, err := u.repo.GetLoginInfoByUsername(username)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password)); err != nil {
		return nil, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
		constants.JwtSubClaim:      user.Id,
		constants.JwtExpClaim:      utils.GetEpochXHoursFromNow(constants.OneDayInHour),
		constants.JwtUsernameClaim: username,
	})
	signedToken, err := token.SignedString([]byte(u.envManager.GetJWTSignKey()))
	if err != nil {
		return nil, err
	}
	return &signedToken, nil
}
