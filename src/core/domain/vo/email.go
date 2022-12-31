package vo

import (
	"fmt"
	"regexp"

	"golang.org/x/xerrors"
)

type Email string

var (
	emailRegExp = regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`)
)

// メールは、英数字記号50文字以下とする。
// RFCに準拠。
const emailMaxLength = 50

func NewEmail(email string) (Email, error) {
	if len(email) == 0 {
		fmt.Printf("メールアドレスは必須入力です。")
		return Email(""), xerrors.New("メールアドレスは必須入力です。")
	}

	if len(email) > emailMaxLength {
		fmt.Printf("メールアドレスを、文字以下で入力してください。現在文字入力されています。")
		return Email(""), xerrors.Errorf("メールアドレスを、%d 文字以下で入力してください。現在%s文字入力されています。", emailMaxLength, email)
	}

	if ok := emailRegExp.MatchString(email); !ok {
		fmt.Printf("入力されたメールアドレスは%sです。フォーマットが正しくありません。", email)
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
