package tmpblogdm

import (
	"unicode/utf8"

	myerror "github.com/schwarzwald0906/My_Portfolio/src/core/myerror"
)

type BlogContent string

const blogContentMaxLength = 10000

func NewBlogContent(blogContent string) (BlogContent, error) {
	if blogContent == "" {
		return BlogContent(""), myerror.BadRequestWrapf("本文は必須入力です。")
	}
	if length := utf8.RuneCountInString(blogContent); length > blogContentMaxLength {
		return BlogContent(""),
			myerror.BadRequestWrapf("本文を、%d文字以下で入力してください。現在%d文字入力されています。", blogContentMaxLength, length)
	}
	return BlogContent(blogContent), nil
}

func (e BlogContent) String() string {
	return string(e)
}

func (e BlogContent) Equals(e2 BlogContent) bool {
	return e.String() == e2.String()
}
