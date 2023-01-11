package vo

import (
	"regexp"

	myerror "github.com/schwarzwald0906/My_Portfolio/src/core/myerror"
	"golang.org/x/crypto/bcrypt"
)

type Password string

var (
	passwordRegExp = []*regexp.Regexp{regexp.MustCompile(`[A-Za-z]`), regexp.MustCompile(`\d`), regexp.MustCompile(`[!-/:-@{-~]`)}
)

func NewPassword(pass string) (Password, error) {
	if pass == "" {
		return "", myerror.BadRequestWrapf("パスワードを入力してください")
	}

	if l := len(pass); l > 31 || l < 9 {
		return "", myerror.BadRequestWrapf("パスワードは、英数字記号8文字以上30文字以下で入力してください。現在%s文字入力されています。", pass)
	}

	if !(regexp.MustCompile("^[0-9a-zA-Z!-/:-@[-`{-~]+$").Match([]byte(pass))) { // 英数字記号以外を使っているか判定
		return "", myerror.BadRequestWrapf("パスワードは、英数字記号8文字以上30文字以下で入力してください。")
	}

	// インスタンスを作成
	for _, r := range passwordRegExp {
		if r.FindString(pass) == "" {
			return "", myerror.BadRequestWrapf("パスワードは、英数字記号をそれぞれ必ず1文字以上含むように入力してください。")
		}
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return "", myerror.InternalServerWrapf("パスワードのハッシュ化に失敗しました")
	}
	return Password(string(hashPass)), nil
}

func (p Password) String() string {
	return string(p)
}
