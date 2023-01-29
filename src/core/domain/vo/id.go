package vo

import (
	"context"

	"github.com/google/uuid"
	myerror "github.com/schwarzwald0906/My_Portfolio/src/core/myerror"
)

type ID string

func NewID() ID {
	return ID(uuid.New().String())
}

func NewIDByStr(c context.Context, id string) (ID, error) {
	if id == "" {
		return "", myerror.BadRequestWrapf("IDは必須入力です。")
	}
	return ID(id), nil
}

func (id ID) Equals(id2 ID) bool {
	return string(id) == string(id2)
}

func (id ID) String() string {
	return string(id)
}
