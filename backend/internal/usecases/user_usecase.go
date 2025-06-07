package usecases

type UserUsecase interface {
	Create()
	Get()
	Update()
	Delete()
}

type UserService struct {
	repo any
}

func NewUserUsercase(repo any) UserUsecase {
	return &UserService{repo: repo}
}

func (u *UserService) CreateUser(req dtos.CreateUserReq) error {
	// TODO: hash password
	// TODO: add password to user entity
	// hashe
	newUser := entities.User{
		Id:        uuid.NewString(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Email:     req.Email,
		HashedPassword: ,
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
