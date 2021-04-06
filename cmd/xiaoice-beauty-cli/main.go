package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/icepie/xiaoice-beauty/client"
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
			fmt.Println("rte: ", rte)
		}
	}

	// from url
	if url != "" {
		rte, err := ib.AnalyzeImgByUrl(url)
		if err != nil {
			fmt.Println("error: ", err)
		} else {
			fmt.Println("rte: ", rte)
		}
	}

}
