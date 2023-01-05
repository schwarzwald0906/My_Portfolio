package vo

import "time"

type CreatedAt time.Time

func NewCreatedAt() (CreatedAt, error) {
	return CreatedAt(time.Now()), nil
}
func NewCreatedAtByVal(val time.Time) (CreatedAt, error) {
	return CreatedAt(val), nil
}
func (e CreatedAt) Time() time.Time {
	return time.Time(e)
}
