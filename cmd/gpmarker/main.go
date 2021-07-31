package main

import (
	"flag"
	"fmt"
	"github.com/gari8/gpmarker"
)

type Content struct {
	FilePath string
	HelpMode bool
}

func main() {
	var content Content
	flag.BoolVar(&content.HelpMode, "H", false, "help-mode")
	flag.StringVar(&content.FilePath, "F", "", "cli-mode file path")
	flag.Parse()
	if content.HelpMode {
		gpmarker.PrintAny(gpmarker.PYellow, helpText)
	} else if content.FilePath != "" {
		jl, err := gpmarker.WalkDirectory(content.FilePath)
		if err != nil {
			gpmarker.PrintAny(gpmarker.PRed, "error: file not found")
		}
		if jl != nil {
			jl.Preview()
		}
	} else {
		jl, err := gpmarker.WalkDirectory(".")
		if err != nil {
			gpmarker.PrintAny(gpmarker.PRed, "error: file not found")
		}
		if jl != nil {
			jl.Preview()
		}
	}
	fmt.Println(guideText)
}

const helpText = `	gpmarker CLI
	options
	-F:
		cli mode
	-H:
		cli guide
`

const guideText = `How to use CLI?:
	gpmarker CLI MODE
	-H: help
	-F: instruct file path`
