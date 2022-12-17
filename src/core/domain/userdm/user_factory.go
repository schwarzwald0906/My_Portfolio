package userdm

import (
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"
)

func GenWhenCreate(email vo.Email, password vo.Password) (*User, error) {
	createdAt, err := vo.NewCreatedAt()
	if err != nil {
		return nil, err
	}

	updatedAt, err := vo.NewUpdatedAt()
	if err != nil {
		return nil, err
	}

	return newUser(UserID(NewUserID()), email, password, createdAt, updatedAt)
}
