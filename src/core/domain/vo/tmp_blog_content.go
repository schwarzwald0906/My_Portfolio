package vo

import (
	"golang.org/x/xerrors"
)

type TmpBlogContent string

// 本文は、30文字以下とする。
const tmpBlogContentMaxLength = 10000

func NewTmpBlogContent(tmpBlogContent string) (TmpBlogContent, error) {
	if len(tmpBlogContent) > tmpBlogContentMaxLength {
		return TmpBlogContent(""), xerrors.Errorf("本文を、%d 文字以下で入力してください。現在%s文字入力されています。", tmpBlogContentMaxLength, tmpBlogContent)
	}
	return TmpBlogContent(tmpBlogContent), nil
}

func (e TmpBlogContent) Value() string {
	return string(e)
}

func (e TmpBlogContent) Equals(e2 TmpBlogContent) bool {
	return e.Value() == e2.Value()
}
