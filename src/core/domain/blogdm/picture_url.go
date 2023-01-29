package tmpblogdm

import (
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"
	myerror "github.com/schwarzwald0906/My_Portfolio/src/core/myerror"
)

type PictureURL vo.URL

func NewPictureURL(urlStr string) (PictureURL, error) {
	if urlStr == "" {
		return PictureURL(""), myerror.BadRequestWrapf("画像がアップロードされていません。")
	}
	url, err := vo.NewURL(urlStr)
	return PictureURL(url), err
}

func (id PictureURL) String() string {
	return string(id)
}

func (id PictureURL) Equals(id2 PictureURL) bool {
	return string(id) == string(id2)
}
