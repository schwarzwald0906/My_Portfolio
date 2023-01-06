package tmpblogdm

import (
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"
)

type TmpBlogID vo.ID

func NewTmpBlogID() TmpBlogID {
	return TmpBlogID(vo.NewID())
}

func NewTmpBlogIDByStr(idStr string) (TmpBlogID, error) {
	id, err := vo.NewIDByStr(idStr)
	return TmpBlogID(id), err
}

func (id TmpBlogID) String() string {
	return string(id)
}

func (id TmpBlogID) Equals(id2 TmpBlogID) bool {
	return string(id) == string(id2)
}
