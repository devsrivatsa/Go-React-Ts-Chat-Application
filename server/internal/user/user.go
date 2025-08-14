package user

import (
	"context"
	"time"
)

type User struct {
	ID        int64     `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreateUserRequest struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type CreateUserResponse struct {
	ID       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	// GetUserByID(ctx context.Context, id int64) (*User, error)
	// GetUserByEmail(ctx context.Context, email string) (*User, error)
	// UpdateUser(ctx context.Context, user *User) (*User, error)
	// DeleteUser(ctx context.Context, id int64) error
}

type Service interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error)
}
