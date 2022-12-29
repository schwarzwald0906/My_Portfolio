package vo

type TmpPicture string

func NewPicture(tmpPicture string) (TmpPicture, error) {
	return TmpPicture(tmpPicture), nil
}

func (e TmpPicture) Value() string {
	return string(e)
}

func (e TmpPicture) Equals(e2 TmpPicture) bool {
	return e.Value() == e2.Value()
}
