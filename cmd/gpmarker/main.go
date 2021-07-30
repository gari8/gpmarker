package main

import (
	"gpmarker"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(gpmarker.Analyzer) }
