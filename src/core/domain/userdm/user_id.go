package userdm

import (
	"github.com/ymdd1/mytweet/src/core/domain/vo"
)

type UserID vo.ID

func NewUserID() vo.ID {
	return vo.NewID()
}

// func NewUserID() UserID {
// 	return UserID(uuid.New().String())
// }

func NewUserIDByStr(idStr string) (UserID, error) {
	id, err := vo.NewIDByStr(idStr)
	return UserID(id), err
}
func (id UserID) Equals(id2 UserID) bool {
	return string(id) == string(id2)
}

func (id UserID) String() string {
	return string(id)
}
