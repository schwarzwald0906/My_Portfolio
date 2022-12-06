package userdm

import "context"

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	FindByUserID(ctx context.Context, userID string) error
	FindByEmailID(ctx context.Context, email string) error
}
