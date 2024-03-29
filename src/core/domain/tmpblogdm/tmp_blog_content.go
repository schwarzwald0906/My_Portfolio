package tmpblogdm

import (
	"unicode/utf8"

	myerror "github.com/schwarzwald0906/My_Portfolio/src/core/myerror"
)

type TmpBlogContent string

const tmpBlogContentMaxLength = 10000

func NewTmpBlogContent(tmpBlogContent string) (TmpBlogContent, error) {
	if length := utf8.RuneCountInString(tmpBlogContent); length > tmpBlogContentMaxLength {
		return TmpBlogContent(""),
			myerror.BadRequestWrapf("本文を、%d文字以下で入力してください。現在%d文字入力されています。", tmpBlogContentMaxLength, length)
	}
	return TmpBlogContent(tmpBlogContent), nil
}

func (e TmpBlogContent) String() string {
	return string(e)
}

func (e TmpBlogContent) Equals(e2 TmpBlogContent) bool {
	return e.String() == e2.String()
}
