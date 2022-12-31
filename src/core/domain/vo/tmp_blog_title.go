package vo

import (
	"golang.org/x/xerrors"
)

type TmpBlogTitle string

// タイトルは、30文字以下とする。
const tmpBlogTitleMaxLength = 30

func NewTmpBlogTitle(tmpBlogTitle string) (TmpBlogTitle, error) {
	if len(tmpBlogTitle) > tmpBlogTitleMaxLength {
		return TmpBlogTitle(""),
			xerrors.Errorf("タイトルを、%d 文字以下で入力してください。現在%s文字入力されています。", tmpBlogTitleMaxLength, tmpBlogTitle)
	}
	return TmpBlogTitle(tmpBlogTitle), nil
}

func (e TmpBlogTitle) Value() string {
	return string(e)
}

func (e TmpBlogTitle) Equals(e2 TmpBlogTitle) bool {
	return e.Value() == e2.Value()
}
