package userdm

import (
	"github.com/ymdd1/mytweet/src/core/domain/vo"
)

type User struct {
	id        vo.UserId
	email     vo.Email
	password  vo.Password
	createdAt vo.CreatedAt
	updatedAt vo.UpdatedAt
}

func NewUser(id vo.UserId, email vo.Email, password vo.Password, createdAt vo.CreatedAt, updatedAt vo.UpdatedAt) (*User, error) {
	return &User{
		id:        id,
		email:     email,
		password:  password,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

func (u *User) ID() vo.UserId {
	return u.id
}
func (u *User) Email() vo.Email {
	return u.email
}
func (u *User) Password() vo.Password {
	return u.password
}
func (u *User) CreatedAt() vo.CreatedAt {
	return u.createdAt
}
func (u *User) UpdatedAt() vo.UpdatedAt {
	return u.updatedAt
}

// func (u *User) Equals(u2 *User) bool {
// 	return u.id.Equals(u2.id)
// }
