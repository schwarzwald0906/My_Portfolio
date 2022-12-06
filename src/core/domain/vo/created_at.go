package vo

import "time"

type CreatedAt time.Time

func NewCreatedAt() CreatedAt {
	return CreatedAt(time.Now())
}
func NewCreatedAtByVal(val time.Time) CreatedAt {
	return CreatedAt(time.Now())
}
func (e CreatedAt) Value() time.Time {
	return time.Time(e)
}
