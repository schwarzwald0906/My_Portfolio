package vo

import "time"

type UpdatedAt time.Time

func NewUpdatedAt() (UpdatedAt, error) {
	return UpdatedAt(time.Now()), nil
}
func NewUpdatedAtByVal(val time.Time) (UpdatedAt, error) {
	return UpdatedAt(val), nil
}
func (e UpdatedAt) Time() time.Time {
	return time.Time(e)
}
