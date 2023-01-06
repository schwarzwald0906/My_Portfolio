package tmpblogdm

import (
	"unicode/utf8"

	"golang.org/x/xerrors"
)

type TmpBlogTitle string

const tmpBlogTitleMaxLength = 30

func NewTmpBlogTitle(tmpBlogTitle string) (TmpBlogTitle, error) {
	if length := utf8.RuneCountInString(tmpBlogTitle); length > tmpBlogTitleMaxLength {
		return TmpBlogTitle(""),
			xerrors.Errorf("タイトルを、%d文字以下で入力してください。現在%d文字入力されています。", tmpBlogTitleMaxLength, length)
	}
	return TmpBlogTitle(tmpBlogTitle), nil
}

func (e TmpBlogTitle) String() string {
	return string(e)
}

func (e TmpBlogTitle) Equals(e2 TmpBlogTitle) bool {
	return e.String() == e2.String()
}
