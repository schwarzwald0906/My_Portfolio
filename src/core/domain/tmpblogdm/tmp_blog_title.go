package tmpblogdm

import (
	"unicode/utf8"

	"golang.org/x/xerrors"
)

type TmpBlogTitle string

const tmpBlogTitleMaxLength = 30

func NewTmpBlogTitle(tmpBlogTitle string) (TmpBlogTitle, error) {
	if utf8.RuneCountInString(tmpBlogTitle) > tmpBlogTitleMaxLength {
		return TmpBlogTitle(""),
			xerrors.Errorf("タイトルを、%d文字以下で入力してください。現在%s文字入力されています。", tmpBlogTitleMaxLength, tmpBlogTitle)
	}
	return TmpBlogTitle(tmpBlogTitle), nil
}

func (e TmpBlogTitle) String() string {
	return string(e)
}

func (e TmpBlogTitle) Equals(e2 TmpBlogTitle) bool {
	return e.String() == e2.String()
}
