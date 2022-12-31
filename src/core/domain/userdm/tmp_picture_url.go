package userdm

import "github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"

type TmpPictureURL vo.URL

func NewTmpPictureURL(urlStr string) (TmpPictureURL, error) {
	url, err := vo.NewURL(urlStr)
	return TmpPictureURL(url), err
}
func (url TmpPictureURL) Equals(url2 TmpPictureURL) bool {
	return string(url) == string(url2)
}

func (url TmpPictureURL) String() string {
	return string(url)
}