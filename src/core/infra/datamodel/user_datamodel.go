package datamodel

import "time"

type User struct {
	id        string
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

func (u *User) ID() string {
	return u.id
}
func (u *User) Email() string {
	return u.email
}
func (u *User) Password() string {
	return u.password
}
func (u *User) CreatedAt() time.Time {
	return time.Time{}
}
func (u *User) UpdatedAt() time.Time {
	return time.Time{}
}
