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

func (u *UserService) Create() {}
func (u *UserService) Get()    {}
func (u *UserService) Update() {}
func (u *UserService) Delete() {}
