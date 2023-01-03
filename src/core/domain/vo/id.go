package vo

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type ID string

func NewID() ID {
	return ID(uuid.New().String())
}

func NewIDByStr(id string) (ID, error) {
	if id == "" {
		return "", xerrors.New("id must be not empty")
	}
	return ID(id), nil
}

func (id ID) Equals(id2 ID) bool {
	return string(id) == string(id2)
}

func (id ID) String() string {
	return string(id)
}
