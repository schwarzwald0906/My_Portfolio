package vo

import "time"

type UpdatedAt time.Time

func NewUpdatedAt() UpdatedAt {
	return UpdatedAt(time.Now())
}
func NewUpdatedAtByVal(val time.Time) UpdatedAt {
	return UpdatedAt(time.Now())
}
func (e UpdatedAt) Value() time.Time {
	return time.Time(e)
}
