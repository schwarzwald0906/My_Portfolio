package userdm

import "context"

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	FindByID(ctx context.Context, userID string) (*User, error)
}
