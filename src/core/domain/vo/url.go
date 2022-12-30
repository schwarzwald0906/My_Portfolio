package vo

import (
	"regexp"

	"golang.org/x/xerrors"
)

var (
	urlRegExp = regexp.MustCompile(`^https?://[\w/:%#\$&\?\(\)~\.=\+\-]+$`)
)

type URL string

func NewURL(url string) URL {
	return URL(url)
}

func NewURLStr(url string) (URL, error) {
	if ok := urlRegExp.MatchString(url); !ok {
		return "", xerrors.New("This is not a URL.")
	}
	return URL(url), nil
}
