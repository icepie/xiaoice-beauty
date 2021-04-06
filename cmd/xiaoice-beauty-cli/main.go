package main

import (
	"fmt"
	"github.com/icepie/xiaoice-beauty/client"
)

func main() {
	ib, err := client.New()
	if err != nil {
		fmt.Println("error: ", err)
	}

	rte, err := ib.AnalyzeImgByUrl("https://goss4.cfp.cn/creative/vcg/800/new/VCG41N860837492.jpg")
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("rte: ", rte)

}
