package vo

import (
	"regexp"

	myerror "github.com/schwarzwald0906/My_Portfolio/src/core/myerror"
)

var (
	urlRegExp = regexp.MustCompile(`^https?://[\w/:%#\$&\?\(\)~\.=\+\-]+$`)
)

type URL string

func NewURL(url string) (URL, error) {
	if ok := urlRegExp.MatchString(url); !ok {
		return "", myerror.BadRequestWrapf("入力されたURLは%sです。フォーマットが正しくありません。", url)
	}
	return URL(url), nil
}

func (url URL) Equals(url2 URL) bool {
	return string(url) == string(url2)
}
