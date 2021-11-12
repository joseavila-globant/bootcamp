package entitities

import "context"

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      int    `json:"age"`
	Password string `json:"password,omitempty"`
}

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	ViewUser(ctx context.Context, id string) (User, error)
	UpdateUser(ctx context.Context, id string, name string, phone string, email string) error
	DeleteUser(ctx context.Context, id string) error
	ListUser(ctx context.Context, limit string, offset string) (string, error)
	AuthenticateUser(ctx context.Context, email string, password string) (string, error)
}
