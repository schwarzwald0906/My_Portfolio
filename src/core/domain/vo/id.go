package vo

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	myerror "github.com/schwarzwald0906/My_Portfolio/src/core/myerror"
)

type ID string

func NewID() ID {
	return ID(uuid.New().String())
}

func NewIDByStr(c *gin.Context, id string) ID {
	var ERR_KEY string
	if id == "" {
		c.Set(ERR_KEY, myerror.BadRequestWrapf("IDは必須入力です。"))
		return ""
	}
	return ID(id)
}

func (id ID) Equals(id2 ID) bool {
	return string(id) == string(id2)
}

func (id ID) String() string {
	return string(id)
}
