package vo

import (
	"regexp"

	"golang.org/x/xerrors"
)

type Email string

var (
	emailFormat = `^[^@]+@[^@]+\.[^@]+$`
	emailRegExp = regexp.MustCompile(emailFormat)
)

// メールは、英数字記号50文字以下とする。
// RFCに準拠。
const emailMaxLength = 50

func NewEmail(email string) (Email, error) {
	if len(email) == 0 {
		return Email(""), xerrors.New("メールアドレスは必須入力です。")
	}

	if len(email) > emailMaxLength {
		return Email(""), xerrors.Errorf("メールアドレスを、%d 文字以下で入力してください。現在%s文字入力されています。", emailMaxLength, email)
	}

	if ok := emailRegExp.MatchString(email); !ok {
		return Email(""), xerrors.Errorf("フォーマットが正しくありません。")
	}

	return Email(email), nil
}

func (e Email) Value() string {
	return string(e)
}

func (e Email) Equals(e2 Email) bool {
	return e.Value() == e2.Value()
}
