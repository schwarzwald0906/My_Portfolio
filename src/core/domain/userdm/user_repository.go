package userdm

import (
	"context"

	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"
)

type UserRepository interface {
	Create(c context.Context, user *User) error
	FindByUserID(c context.Context, userId UserID) (*User, error)
	FindByEmailID(c context.Context, email vo.Email) (*User, error)
}
