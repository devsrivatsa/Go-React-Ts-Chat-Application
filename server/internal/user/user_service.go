package user

import (
	"context"
	"strconv"
	"time"

	"github.com/devsrivatsa/chat_app_go-ts-react/utils"
	jwt "github.com/golang-jwt/jwt/v5"
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

type JwtCustomClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *UserService) Login(c context.Context, req *LoginUserRequest) (*LoginUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginUserResponse{}, err
	}

	if err := utils.CheckPassword(req.Password, user.Password); err != nil {
		return &LoginUserResponse{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaims{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    strconv.Itoa(int(user.ID)),
		},
	})
	accessToken, err := token.SignedString([]byte("some_secret"))
	if err != nil {
		return &LoginUserResponse{}, err
	}

	return &LoginUserResponse{
		accessToken: accessToken,
		ID:          strconv.Itoa(int(user.ID)),
		Username:    user.Username,
		Email:       user.Email,
	}, nil
}
