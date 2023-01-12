package vo

import (
	"regexp"

	"github.com/gin-gonic/gin"
	myerror "github.com/schwarzwald0906/My_Portfolio/src/core/myerror"
	"golang.org/x/crypto/bcrypt"
)

type Password string

var (
	passwordRegExp = []*regexp.Regexp{regexp.MustCompile(`[A-Za-z]`), regexp.MustCompile(`\d`), regexp.MustCompile(`[!-/:-@{-~]`)}
)

func NewPassword(c *gin.Context, pass string) Password {
	var ERR_KEY string
	if pass == "" {
		c.Set(ERR_KEY, myerror.BadRequestWrapf("パスワードを入力してください"))
		return ""
	}

	if l := len(pass); l > 31 || l < 9 {
		c.Set(ERR_KEY, myerror.BadRequestWrapf("パスワードは、英数字記号8文字以上30文字以下で入力してください。現在%s文字入力されています。", pass))
		return ""
	}

	if !(regexp.MustCompile("^[0-9a-zA-Z!-/:-@[-`{-~]+$").Match([]byte(pass))) { // 英数字記号以外を使っているか判定
		c.Set(ERR_KEY, myerror.BadRequestWrapf("パスワードは、英数字記号8文字以上30文字以下で入力してください。"))
		return ""
	}

	// インスタンスを作成
	for _, r := range passwordRegExp {
		if r.FindString(pass) == "" {
			c.Set(ERR_KEY, myerror.BadRequestWrapf("パスワードは、英数字記号をそれぞれ必ず1文字以上含むように入力してください。"))
			return ""
		}
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		c.Set(ERR_KEY, myerror.InternalServerWrapf("パスワードのハッシュ化に失敗しました"))
		return ""
	}
	return Password(string(hashPass))
}

func (p Password) String() string {
	return string(p)
}
