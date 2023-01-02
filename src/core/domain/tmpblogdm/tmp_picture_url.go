package tmpblogdm

import "github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"

type TmpPictureURL vo.URL

func NewTmpPictureURL(urlStr string) (TmpPictureURL, error) {
	url, err := vo.NewURL(urlStr)
	return TmpPictureURL(url), err
}

func (id TmpPictureURL) String() string {
	return string(id)
}

func (id TmpPictureURL) Equals(id2 TmpPictureURL) bool {
	return string(id) == string(id2)
}
