package vo

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"
)

type Password string

var (
	passwordRegExp = []*regexp.Regexp{regexp.MustCompile(`[A-Za-z]`), regexp.MustCompile(`\d`), regexp.MustCompile(`[!-/:-@{-~]`)}
)

// パスワードは、英数字記号8文字以上30文字以下とする。
// 必ず英字数字記号を1字以上使わなければならないとする。
func NewPassword(pass string) (Password, error) {
	if pass == "" {
		return "", xerrors.New("パスワードを入力してください")
	}

	if l := len(pass); l > 31 || l < 9 {
		return "", xerrors.Errorf("パスワードは、英数字記号8文字以上30文字以下で入力してください。現在%s文字入力されています。", pass)
	}

	if !(regexp.MustCompile("^[0-9a-zA-Z!-/:-@[-`{-~]+$").Match([]byte(pass))) { // 英数字記号以外を使っているか判定
		return "", xerrors.New("パスワードは、英数字記号8文字以上30文字以下で入力してください。")
	}

	// インスタンスを作成
	for _, r := range passwordRegExp {
		if r.FindString(pass) == "" {
			return "", xerrors.New("パスワードは、英数字記号をそれぞれ必ず1文字以上含むように入力してください。")
		}
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return "", xerrors.New("パスワードのハッシュ化に失敗しました")
	}
	return Password(string(hashPass)), nil
}

func (p Password) Value() string {
	return string(p)
}
