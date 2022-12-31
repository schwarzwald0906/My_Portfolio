package userdm

import (
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"
)

type BlogID vo.ID

func NewBlogID() vo.ID {
	return vo.NewID()
}

func NewBlogIDByStr(idStr string) (BlogID, error) {
	id, err := vo.NewIDByStr(idStr)
	return BlogID(id), err
}
func (id BlogID) Equals(id2 BlogID) bool {
	return string(id) == string(id2)
}

func (id BlogID) String() string {
	return string(id)
}
