package userdm

import (
	"github.com/ymdd1/mytweet/src/core/domain/vo"
)

type User struct {
	id         UserID
	email      vo.Email
	password   vo.Password
	created_at vo.Created_at
	updated_at vo.Updated_at
}

func NewUser(id UserID, email vo.Email, password vo.Password, created_at vo.Created_at, updated_at vo.Updated_at) (*User, error) {
	return &User{
		id:         id,
		email:      email,
		password:   password,
		created_at: created_at,
		updated_at: updated_at,
	}, nil
}

func (u *User) ID() UserID {
	return u.id
}
func (u *User) Email() vo.Email {
	return u.email
}
func (u *User) Password() vo.Password {
	return u.password
}
func (u *User) Created_at() vo.Created_at {
	return u.created_at
}
func (u *User) Updated_at() vo.Updated_at {
	return u.updated_at
}
func (u *User) Equals(u2 *User) bool {
	return u.id.Equals(u2.id)
}
