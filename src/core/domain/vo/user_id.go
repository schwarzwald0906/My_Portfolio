package vo

import "github.com/google/uuid"

type UserId string

// SQLから連番取得みたいな処理が必要だが、未実装
func NewUserID() UserId {
	uuidObj, _ := uuid.NewUUID()
	return UserId(uuidObj.String())
}

func (e UserId) Value() string {
	return string(e)
}
