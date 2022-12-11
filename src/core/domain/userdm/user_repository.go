package userdm

import (
	"context"

	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByUserID(ctx context.Context, userId UserID) (*User, error)
	FindByEmailID(ctx context.Context, email vo.Email) (*User, error)
}
