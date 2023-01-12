package tmpblogdm

import (
	"github.com/gin-gonic/gin"
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"
)

type BlogID vo.ID

func NewBlogID() BlogID {
	return BlogID(vo.NewID())
}

func NewBlogIDByStr(c *gin.Context, idStr string) BlogID {
	id := vo.NewIDByStr(c, idStr)
	return BlogID(id)
}

func (id BlogID) String() string {
	return string(id)
}

func (id BlogID) Equals(id2 BlogID) bool {
	return string(id) == string(id2)
}
