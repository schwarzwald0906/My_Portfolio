package userdm

import (
	"context"

	"github.com/ymdd1/mytweet/src/core/domain/vo"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByUserID(ctx context.Context, userId vo.UserId) (*User, error)
	FindByEmailID(ctx context.Context, email vo.Email) (*User, error)
}
