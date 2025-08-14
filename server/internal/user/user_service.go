package user

import (
	"context"
	"strconv"
	"time"

	"github.com/devsrivatsa/chat_app_go-ts-react/utils"
)

type UserService struct {
	repository Repository
	timeout    time.Duration
}

func NewUserService(repository Repository) *UserService {
	return &UserService{
		repository: repository,
		timeout:    time.Duration(2 * time.Second),
	}
}

func (s *UserService) CreateUser(c context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	user, err := s.repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return &CreateUserResponse{ID: strconv.Itoa(int(user.ID)), Username: user.Username, Email: user.Email}, nil
}
