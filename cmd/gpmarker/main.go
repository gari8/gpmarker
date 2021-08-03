package main

import (
	"flag"
	"fmt"
	"github.com/gari8/gpmarker"
)

type Content struct {
	FilePath string
	HelpMode bool
	RequireRows bool
}

func main() {
	var content Content
	flag.BoolVar(&content.RequireRows, "r", false, "require code rows default false")
	flag.BoolVar(&content.HelpMode, "h", false, "help-mode default false")
	flag.StringVar(&content.FilePath, "p", ".", "cli-mode file path")
	flag.Parse()
	if content.HelpMode {
		gpmarker.PrintAny(gpmarker.PYellow, helpText)
	} else {
		jl, err := gpmarker.WalkDirectory(content.FilePath)
		if err != nil {
			gpmarker.PrintAny(gpmarker.PRed, "error: file not found")
		}
		if jl != nil {
			jl.Preview(content.RequireRows)
		}
	}
	fmt.Println(guideText)
}

const helpText = `	gpmarker CLI
	options
	-p:
		you can instruct file path
	-h:
		show help message
	-r:
		show source code near comment
`

const guideText = `How to use CLI?:
	gpmarker CLI MODE
	-h: show help message
	-p: you can instruct file path
	-r: show source code near comment`
