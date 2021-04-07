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

### 使用

> 注意将 `GOPTAH` 添加到环境变量

#### 帮助

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

#### 范例

> 那我就拿自己的照片鉴定一下颜值吧 (逃

```bash

$ xiaoice-beauty-cli -f me.png

## 鉴定结果

        - 得分:  9.3

        - 简述: 要是谁能有吴彦祖这样的男朋友，做梦都会笑醒吧？（从网络上收集到的吴彦祖照片统计显示，吴彦祖的颜值均值在9.3分上下）

        - 分析图: https://mediaplatform.xiaoice.com/image/fetchimage?key=******

## 人物信息

        - 性别: male

        - 是否为名人: True

        - 是否为表情包: False

        - 面部像素区域: [[(99, 85), (211, 85), (99, 197), (211, 197)]]

## 评分细则

        - 韩国90后女性:  9.3

        - 法国90后女性:  8.7

        - 法国90后男性:  8.3

```