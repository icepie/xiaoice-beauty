package client

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"xiaoice-beauty/model"
)

// UploadImg 通过 base64 上传图片
func (ib IceBeauty) UploadImg(base64 string) (string, error) {

	body := strings.NewReader(base64)

	req, err := http.NewRequest("POST", UploadUrl, body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")

	// 设置 cookies
	for _, cookie := range ib.Cookies {
		req.AddCookie(cookie)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 将数据流转换为 []byte
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tmp model.UploadImgRte

	err = json.Unmarshal([]byte(string(b)), &tmp)
	if err != nil {
		return "", err
	}

	fullUrl := tmp.Host + tmp.URL

	return fullUrl, nil
}

// UploadImgByFile 通过文件上传图片
func (ib IceBeauty) UploadImgByFile(filePath string) (string, error) {

	srcByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	res := base64.StdEncoding.EncodeToString(srcByte)

	return ib.UploadImg(res)
}

// UploadImgByFile 通过链接上传图片
func (ib IceBeauty) UploadImgByUrl(imageUrl string) (string, error) {

	resp, err := http.Get(imageUrl)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	// 读取获取的[]byte数据
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	res := base64.StdEncoding.EncodeToString(data)

	return ib.UploadImg(res)
}
