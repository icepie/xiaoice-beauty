package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/icepie/xiaoice-beauty/client"
	"github.com/icepie/xiaoice-beauty/model"
)

var (
	file, url string
	ishelp    bool
)

func printHelp() {
	fmt.Fprintf(os.Stderr, `
xiaoice-beauty-cli
	A command line tools for face recognition and rating, based on Microsoft XiaoIce API

usage:
	xiaoice-beauty-cli [-h] [-f <FILE>] [-u <URL>]

options:
`)
	flag.PrintDefaults()
}

func printResult(rte model.AnalyzeImgRte) {
	fmt.Printf(`
## 鉴定结果

	- 得分: %.1f

	- 简述: %s

	- 分析图: %s

## 人物信息

	- 性别: %s

	- 是否为名人: %s

	- 是否为表情包: %s

	- 面部像素区域: %s

## 评分细则

	- %s: %.1f
	
	- %s: %.1f

	- %s: %.1f

`, rte.Content.Metadata.Score, rte.Content.Text, rte.Content.Metadata.Reportimgurl, rte.Content.Metadata.Gender, rte.Content.Metadata.Isceleb, rte.Content.Metadata.Isemoji, rte.Content.Metadata.FacePoints, rte.Content.Metadata.FbrKey0, rte.Content.Metadata.FbrScore0, rte.Content.Metadata.FbrKey1, rte.Content.Metadata.FbrScore1, rte.Content.Metadata.FbrKey2, rte.Content.Metadata.FbrScore2)
}

func main() {

	flag.StringVar(&file, "f", "", "image from file")
	flag.StringVar(&url, "u", "", "image from url")
	flag.BoolVar(&ishelp, "h", false, "show this help message and exit")

	flag.Parse()

	if ishelp {
		printHelp()
		os.Exit(0)
	}

	if file == "" && url == "" {
		printHelp()
		os.Exit(0)
	} else if file != "" && url != "" {
		fmt.Println("error: Please don't use file and url at the same time!")
		printHelp()
		os.Exit(1)
	}

	// init
	ib, err := client.New()
	if err != nil {
		fmt.Println("error: ", err)
	}

	// from file
	if file != "" {
		rte, err := ib.AnalyzeImgByFile(file)
		if err != nil {
			fmt.Println("error: ", err)
		} else {
			printResult(rte)
		}
	}

	// from url
	if url != "" {
		rte, err := ib.AnalyzeImgByUrl(url)
		if err != nil {
			fmt.Println("error: ", err)
		} else {
			printResult(rte)
		}
	}

}
