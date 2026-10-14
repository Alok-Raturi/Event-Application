package user

import (
	"strings"
	"time"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUsers() ([]User, error) {
	return s.userRepo.GetUsers()
}

func (s *UserService) CreateUser(userRequest *UserRequest) error {

	userRequest.Email = strings.TrimSpace(userRequest.Email)
	userRequest.Name = strings.TrimSpace(userRequest.Name)

	if userRequest.Name == "" || userRequest.Email == "" {
		return ErrInvalidInput
	}

	if userRequest.Password == "" || len(userRequest.Password) < 6 || len(userRequest.Password) > 72 {
		return ErrInvalidInput
	}

	existingUser, err := s.userRepo.GetUserByEmail(userRequest.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return ErrEmailAlreadyExists
	}
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	hashedPassword, err := hashPassword(userRequest.Password)
	if err != nil {
		return err
	}

	user := &User{
		ID:           id,
		Name:         userRequest.Name,
		Email:        userRequest.Email,
		PasswordHash: hashedPassword,
		CreatedAt:    time.Now(),
	}
	err = s.userRepo.CreateUser(user)
	return err
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
