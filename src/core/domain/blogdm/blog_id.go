package tmpblogdm

import (
	"github.com/google/uuid"
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"
)

type BlogID vo.ID

func NewBlogID() BlogID {
	return BlogID(uuid.New().String())
}

func NewBlogIDByStr(idStr string) (BlogID, error) {
	id, err := vo.NewIDByStr(idStr)
	return BlogID(id), err
}

func (id BlogID) String() string {
	return string(id)
}

func (id BlogID) Equals(id2 BlogID) bool {
	return string(id) == string(id2)
}
