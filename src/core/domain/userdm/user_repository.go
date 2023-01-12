package userdm

import (
	"github.com/gin-gonic/gin"
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"
)

type UserRepository interface {
	Create(c *gin.Context, user *User) error
	FindByUserID(c *gin.Context, userId UserID) (*User, error)
	FindByEmailID(c *gin.Context, email vo.Email) (*User, error)
}
