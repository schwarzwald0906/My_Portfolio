package userdm

import (
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"
)

type TmpBlogID vo.ID

func NewTmpBlogID() vo.ID {
	return vo.NewID()
}

func NewTmpBlogIDByStr(idStr string) (TmpBlogID, error) {
	id, err := vo.NewIDByStr(idStr)
	return TmpBlogID(id), err
}
