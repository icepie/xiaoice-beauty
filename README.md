# xiaoice-beauty
今天天气暖暖的, 可我的心却是冰冰的~ (封装了一下微软小冰颜值鉴定接口)

## 封装库

### 获取

```bash
$ go get github.com/icepie/xiaoice-beauty
```

### 调用

```go
package main

import (
	"fmt"
	"github.com/icepie/xiaoice-beauty/client"
)

func main() {
    // 声明对象
	ib, err := client.New()
	if err != nil {
		fmt.Println("error: ", err)
	}
    
	// 从图片链接分析
	rte, err := ib.AnalyzeImgByUrl("https://goss4.cfp.cn/creative/vcg/800/new/VCG41N860837492.jpg")
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("rte: ", rte)
    
    // 从图片文件分析
	rte, err = ib.AnalyzeImgByFile("./a.png")
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("rte: ", rte)
}

``` 

## 命令行

### 获取

```bash
$ go install github.com/icepie/xiaoice-beauty/cmd/xiaoice-beauty-cli
```

### 调用

```bash
$ xiaoice-beauty-cli -h

xiaoice-beauty-cli
	A command line tools for face recognition and rating, based on Microsoft XiaoIce API

usage:
	xiaoice-beauty-cli [-h] [-f <FILE>] [-u <URL>]

options:
  -f string
    	image from file
  -u string
    	image from url
  -h	show this help message and exit

```

