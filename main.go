package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetCooikes() []*http.Cookie {

	resp, err := http.Head("https://ux.xiaoice.com/beautyv3")
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	return resp.Cookies()

}

func UploadIMG(cookies []*http.Cookie, file string) string {

	type Rte struct {
		Host string `json:"Host"`
		URL  string `json:"Url"`
	}

	srcByte, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	res := base64.StdEncoding.EncodeToString(srcByte)

	body := strings.NewReader(res)

	req, err := http.NewRequest("POST", "https://ux.xiaoice.com/api/image/UploadBase64?exp=0", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	// 将数据流转换为 []byte
	b, err := ioutil.ReadAll(resp.Body)

	var tmp Rte
	err = json.Unmarshal([]byte(string(b)), &tmp) //第二个参数要地址传递
	if err != nil {
		fmt.Println("err = ", err)
	}

	return tmp.Host + tmp.URL
}

func main() {

	// 设置cookies
	cookies := GetCooikes()

	a := UploadIMG(cookies, "./VCG41N860837492.jpg")

	println(a)

	type Content struct {
		Imageurl string `json:"imageUrl"`
	}

	type Payload struct {
		Msgid      int64   `json:"MsgId"`
		Createtime int     `json:"CreateTime"`
		Traceid    string  `json:"TraceId"`
		Content    Content `json:"Content"`
	}

	data := Payload{
		// Msgid:      1617684718400,
		// Createtime: 1617684718,
		// Traceid: "3493264258b5726947b0de4eab35d70d",
		Content: Content{Imageurl: a},
		// fill struct
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://ux.xiaoice.com/api/imageAnalyze/Process?service=beauty", body)
	if err != nil {
		// handle err
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	// 将数据流转换为 []byte
	b, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(b))
}
