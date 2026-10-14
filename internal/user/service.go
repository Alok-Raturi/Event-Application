package user

import (
	"errors"
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
	users, err := s.userRepo.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) CreateUser(userRequest *UserRequest) error {
	if strings.TrimSpace(userRequest.Name) == "" || strings.TrimSpace(userRequest.Email) == "" {
		return errors.New("name and email cannot be empty")
	}

	if userRequest.Password == "" || len(userRequest.Password) < 6 || len(userRequest.Password) > 72 {
		return errors.New("password must be between 6 and 72 characters")
	}

	userRequest.Email = strings.TrimSpace(userRequest.Email)
	userRequest.Name = strings.TrimSpace(userRequest.Name)

	existingUser, err := s.userRepo.GetUserByEmail(userRequest.Email)
	if err != nil {
		return errors.New("error checking existing user: " + err.Error())
	}
	if existingUser != nil {
		return errors.New("user with this email already exists")
	}
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	hashedPassword, err := hashpassword(userRequest.Password)
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

func hashpassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
