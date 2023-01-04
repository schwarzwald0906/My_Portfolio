package tmpblogdm

import (
	"unicode/utf8"

	"golang.org/x/xerrors"
)

type BlogContent string

const blogContentMaxLength = 10000

func NewBlogContent(blogContent string) (BlogContent, error) {
	if len(blogContent) == 0 {
		return BlogContent(""), xerrors.New("本文は必須入力です。")
	}
	if utf8.RuneCountInString(blogContent) > blogContentMaxLength {
		return BlogContent(""),
			xerrors.Errorf("本文を、%d文字以下で入力してください。現在%d文字入力されています。", blogContentMaxLength, utf8.RuneCountInString(blogContent))
	}
	return BlogContent(blogContent), nil
}

func (e BlogContent) Value() string {
	return string(e)
}

func (e BlogContent) Equals(e2 BlogContent) bool {
	return e.Value() == e2.Value()
}
