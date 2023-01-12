package userdm

import (
	"github.com/gin-gonic/gin"
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"
)

type UserID vo.ID

func NewUserID() vo.ID {
	return vo.NewID()
}

func NewUserIDByStr(c *gin.Context, idStr string) UserID {
	id := vo.NewIDByStr(c, idStr)
	return UserID(id)
}

func (id UserID) String() string {
	return string(id)
}

func (id UserID) Equals(id2 UserID) bool {
	return string(id) == string(id2)
}
