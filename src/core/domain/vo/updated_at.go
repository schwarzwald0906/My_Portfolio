package vo

type Updated_at string

func NewUpdated_at(updated_at string) (Updated_at, error) {
	return Updated_at(updated_at), nil
}

func (e Updated_at) Value() string {
	return string(e)
}
