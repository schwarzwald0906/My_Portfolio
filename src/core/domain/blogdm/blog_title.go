package tmpblogdm

import (
	"unicode/utf8"

	"golang.org/x/xerrors"
)

type BlogTitle string

const blogTitleMaxLength = 30

func NewBlogTitle(blogTitle string) (BlogTitle, error) {
	if blogTitle == "" {
		return BlogTitle(""), xerrors.New("タイトルは必須入力です。")
	}
	if utf8.RuneCountInString(blogTitle) > blogTitleMaxLength {
		return BlogTitle(""),
			xerrors.Errorf("タイトルを、%d文字以下で入力してください。現在%s文字入力されています。", blogTitleMaxLength, blogTitle)
	}
	return BlogTitle(blogTitle), nil
}

func (e BlogTitle) String() string {
	return string(e)
}

func (e BlogTitle) Equals(e2 BlogTitle) bool {
	return e.String() == e2.String()
}
