package client

import (
	"net/http"
)

const (
	// MainUrl 小冰颜值鉴定主页
	MainUrl = "https://ux.xiaoice.com/beautyv3"
	// UploadUrl 上传图片接口
	UploadUrl = "https://ux.xiaoice.com/api/image/UploadBase64?exp=0"
	// AnalyzeUrl 分析图片接口
	AnalyzeUrl = "https://ux.xiaoice.com/api/imageAnalyze/Process?service=beauty"
)

func getCooikes() ([]*http.Cookie, error) {

	resp, err := http.Head(MainUrl)
	if err != nil {
		return resp.Cookies(), err
	}

	defer resp.Body.Close()

	return resp.Cookies(), nil
}

// IceBeauty 声明基本结构体
type IceBeauty struct {
	Cookies []*http.Cookie
}

// New 生成初始化的 IceBeauty
func New() (ib IceBeauty, err error) {
	ib.Cookies, err = getCooikes()
	if err != nil {
		return ib, err
	}

	return ib, nil
}
