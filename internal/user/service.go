package user

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUsers() ([]User, error) {
	users, err := s.userRepo.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) CreateUser(user *User) error {
	err := s.userRepo.CreateUser(user)
	return err
}
