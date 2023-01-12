package vo

import (
	"regexp"

	"github.com/gin-gonic/gin"
	myerror "github.com/schwarzwald0906/My_Portfolio/src/core/myerror"
)

type Email string

var (
	emailRegExp = regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`)
)

// RFCに準拠。
const emailMaxLength = 50

func NewEmail(c *gin.Context, email string) Email {
	var ERR_KEY string
	if email == "" {
		c.Set(ERR_KEY, myerror.BadRequestWrapf("メールアドレスは必須入力です。"))
		return Email("")
	}

	if len(email) > emailMaxLength {
		c.Set(ERR_KEY, myerror.BadRequestWrapf("メールアドレスを、%d文字以下で入力してください。現在%s文字入力されています。", emailMaxLength, email))
		return Email("")
	}

	if ok := emailRegExp.MatchString(email); !ok {
		c.Set(ERR_KEY, myerror.BadRequestWrapf(" 入力値は、%sです。フォーマットが正しくありません。", email))
		return Email("")
	}

	return Email(email)
}

func (e Email) String() string {
	return string(e)
}

func (e Email) Equals(e2 Email) bool {
	return e.String() == e2.String()
}
