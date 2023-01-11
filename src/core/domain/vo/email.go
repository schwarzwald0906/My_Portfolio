package vo

import (
	"regexp"

	myerror "github.com/schwarzwald0906/My_Portfolio/src/core/myerror"
	"golang.org/x/xerrors"
)

type Email string

var (
	emailRegExp = regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`)
)

// RFCに準拠。
const emailMaxLength = 50

func NewEmail(email string) (Email, error) {

	if email == "" {
		return Email(""), myerror.BadRequestWrapf("メールアドレスは必須入力です。")
	}
	// if email == "" {
	// 	return Email(""), xerrors.New("メールアドレスは必須入力です。")
	// }

	if len(email) > emailMaxLength {
		return Email(""), xerrors.Errorf("メールアドレスを、%d文字以下で入力してください。現在%s文字入力されています。", emailMaxLength, email)
	}

	if ok := emailRegExp.MatchString(email); !ok {
		return Email(""), xerrors.Errorf("フォーマットが正しくありません。")
	}

	return Email(email), nil
}

func (e Email) String() string {
	return string(e)
}

func (e Email) Equals(e2 Email) bool {
	return e.String() == e2.String()
}
