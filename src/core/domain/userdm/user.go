package userdm

import (
	"github.com/ymdd1/mytweet/src/core/domain/vo"
)

type User struct {
	id       UserID
	email    string
	password vo.Password
}

const (
	maxUserNameLength = 20
)

func NewUser(id UserID, email string, password vo.Password) (*User, error) {
	return &User{
		id:       id,
		email:    email,
		password: password,
	}, nil
}

func (u *User) ID() UserID {
	return u.id
}
func (u *User) Email() string {
	return u.email
}
func (u *User) Password() vo.Password {
	return u.password
}
func (u *User) Equals(u2 *User) bool {
	return u.id.Equals(u2.id)
}
