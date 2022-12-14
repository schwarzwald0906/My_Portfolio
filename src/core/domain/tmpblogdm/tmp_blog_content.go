package tmpblogdm

import (
	"unicode/utf8"

	"golang.org/x/xerrors"
)

type TmpBlogContent string

const tmpBlogContentMaxLength = 10000

func NewTmpBlogContent(tmpBlogContent string) (TmpBlogContent, error) {
	if utf8.RuneCountInString(tmpBlogContent) > tmpBlogContentMaxLength {
		return TmpBlogContent(""),
			xerrors.Errorf("本文を、%d文字以下で入力してください。現在%d文字入力されています。", tmpBlogContentMaxLength, utf8.RuneCountInString(tmpBlogContent))
	}
	return TmpBlogContent(tmpBlogContent), nil
}

func (e TmpBlogContent) Value() string {
	return string(e)
}

func (e TmpBlogContent) Equals(e2 TmpBlogContent) bool {
	return e.Value() == e2.Value()
}
