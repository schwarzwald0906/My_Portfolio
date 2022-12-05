package vo

type Created_at string

func NewCreated_at(created_at string) (Created_at, error) {
	return Created_at(created_at), nil
}

func (e Created_at) Value() string {
	return string(e)
}
