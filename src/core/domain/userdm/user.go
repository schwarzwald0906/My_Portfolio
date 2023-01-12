package userdm

import (
	"time"

	"github.com/gin-gonic/gin"
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

func Reconstruct(c *gin.Context, id string, email string, password string, createdAt time.Time, upadatedAt time.Time) (*User, error) {
	var user *User

	newId := NewUserIDByStr(c, id)

	newEmail := vo.NewEmail(c, email)

	newPassword := vo.NewPassword(c, password)

	newCreatedAt, err := vo.NewCreatedAtByVal(createdAt)
	if err != nil {
		return user, err
	}
	newUpadatedAt, err := vo.NewUpdatedAtByVal(upadatedAt)
	if err != nil {
		return user, err
	}

	return &User{
		id:        newId,
		email:     newEmail,
		password:  newPassword,
		createdAt: newCreatedAt,
		updatedAt: newUpadatedAt,
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

func (u *User) Equals(u2 *User) bool {
	return u.id.Equals(u2.id)
}
