package tmpblogdm

import (
	"unicode/utf8"

	myerror "github.com/schwarzwald0906/My_Portfolio/src/core/myerror"
)

type TmpBlogTitle string

const tmpBlogTitleMaxLength = 30

func NewTmpBlogTitle(tmpBlogTitle string) (TmpBlogTitle, error) {
	if length := utf8.RuneCountInString(tmpBlogTitle); length > tmpBlogTitleMaxLength {
		return TmpBlogTitle(""),
			myerror.BadRequestWrapf("タイトルを、%d文字以下で入力してください。現在%d文字入力されています。", tmpBlogTitleMaxLength, length)
	}
	return TmpBlogTitle(tmpBlogTitle), nil
}

func (e TmpBlogTitle) String() string {
	return string(e)
}

func (e TmpBlogTitle) Equals(e2 TmpBlogTitle) bool {
	return e.String() == e2.String()
}
