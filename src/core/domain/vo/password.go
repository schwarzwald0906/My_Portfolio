package vo

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"
)

type Password string

// パスワードは、英数字記号8文字以上30文字以下とする。
//
//	※余裕があれば、必ず大文字小文字数字を使わなければならないとする。→未実装
func NewPassword(pass string) (Password, error) {
	if pass == "" {
		return "", xerrors.New("password must be not empty")
	}

	if l := len(pass); l > 31 || l < 9 {
		return "", xerrors.New("password must be from 8 to 30 characters")
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return "", xerrors.New("can't generate hash password")
	}
	return Password(string(hashPass)), nil
}

func (p Password) Value() string {
	return string(p)
}
