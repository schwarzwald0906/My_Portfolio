package userdm

import (
	"time"

	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"
)

type User struct {
	id        UserID
	email     vo.Email
	password  vo.Password
	createdAt vo.CreatedAt
	updatedAt vo.UpdatedAt
}

func newUser(id UserID, email vo.Email, password vo.Password, createdAt vo.CreatedAt, updatedAt vo.UpdatedAt) (*User, error) {
	return &User{
		id:        id,
		email:     email,
		password:  password,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

func Reconstruct(id string, email string, password string, createdat time.Time, upadatedat time.Time) (*User, error) {
	var user *User

	newid, err := NewUserIDByStr(id)
	if err != nil {
		return user, err
	}
	newemail, err := vo.NewEmail(email)
	if err != nil {
		return user, err
	}
	newpassword, err := vo.NewPassword(password)
	if err != nil {
		return user, err
	}
	newcreatedat, err := vo.NewCreatedAtByVal(createdat)
	if err != nil {
		return user, err
	}
	newupadatedat, err := vo.NewUpdatedAtByVal(upadatedat)
	if err != nil {
		return user, err
	}

	return &User{
		id:        newid,
		email:     newemail,
		password:  newpassword,
		createdAt: newcreatedat,
		updatedAt: newupadatedat,
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
func (u *User) CreatedAt() vo.CreatedAt {
	return u.createdAt
}
func (u *User) UpdatedAt() vo.UpdatedAt {
	return u.updatedAt
}

// func (u *User) Equals(u2 *User) bool {
// 	return u.id.Equals(u2.id)
// }
