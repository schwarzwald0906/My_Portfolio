package tmpblogdm

import (
	"github.com/gin-gonic/gin"
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"
)

type TmpBlogID vo.ID

func NewTmpBlogID() TmpBlogID {
	return TmpBlogID(vo.NewID())
}

func NewTmpBlogIDByStr(c *gin.Context, idStr string) TmpBlogID {
	id := vo.NewIDByStr(c, idStr)
	return TmpBlogID(id)
}

func (id TmpBlogID) String() string {
	return string(id)
}

func (id TmpBlogID) Equals(id2 TmpBlogID) bool {
	return string(id) == string(id2)
}
