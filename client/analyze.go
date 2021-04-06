package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/icepie/xiaoice-beauty/model"
)

// AnalyzeImg 分析颜值
func (ib IceBeauty) AnalyzeImg(imageUrl string) (model.AnalyzeImgRte, error) {

	var rte model.AnalyzeImgRte

	// 获取当前时间戳
	timeUnix := time.Now().Unix()

	// // 生成 Msgid
	// msgId, _ := strconv.Atoi(int(fmt.Fprintf("%d%d", timeUnix, 400)))

	data := model.AnalyzeImgBody{
		//Msgid:      msgId,
		Createtime: timeUnix,
		// Traceid: "3493264258b5726947b0de4eab35d70d",
		Content: model.Content{Imageurl: imageUrl},
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return rte, err
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", AnalyzeUrl, body)
	if err != nil {
		return rte, err
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")

	// 设置 cookies
	for _, cookie := range ib.Cookies {
		req.AddCookie(cookie)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return rte, err
	}

	defer resp.Body.Close()

	// 将数据流转换为 []byte
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return rte, err
	}

	err = json.Unmarshal([]byte(string(b)), &rte)
	if err != nil {
		return rte, err
	}

	return rte, nil
}

// AnalyzeImgByFile 通过文件图片分析颜值
func (ib IceBeauty) AnalyzeImgByFile(filePath string) (rte model.AnalyzeImgRte, err error) {
	url, err := ib.UploadImgByFile(filePath)
	if err != nil {
		return rte, err
	}

	return ib.AnalyzeImg(url)
}

// AnalyzeImgByFile 通过其他图片链接地址分析颜值
func (ib IceBeauty) AnalyzeImgByUrl(imageUrl string) (rte model.AnalyzeImgRte, err error) {
	url, err := ib.UploadImgByUrl(imageUrl)
	if err != nil {
		return rte, err
	}

	return ib.AnalyzeImg(url)
}
