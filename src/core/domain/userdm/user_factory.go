package userdm

import (
	"github.com/ymdd1/mytweet/src/core/domain/vo"
)

func GenWhenCreate(email vo.Email, password vo.Password) (*User, error) {
	return newUser(NewUserID(), email, password, vo.NewCreatedAt(), vo.NewUpdatedAt())
}
