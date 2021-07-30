package gpmarker

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"strings"
)

const doc = "gpmarker is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "gpmarker",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	fset := token.NewFileSet()
	var journalList JournalList
	for _, file := range pass.Files {
		cMap := ast.NewCommentMap(fset, file, file.Comments)
		for _, c := range cMap.Comments() {
			mark, tp, text := splitSentence(c.Text())
			if mark == "mark" {
				journal := Journal{
					Type: tp,
					FileName: file.Name.String(),
					Pos: int(c.Pos()),
					Text: text,
				}
				journalList.Add(&journal)
			}
		}
	}

	journalList.Preview()
	return nil, nil
}

// mark, type, text
func splitSentence(s string) (string, string, string) {
	sentences := arrangeLine(strings.Split(s, " "))
	mt := arrangeLine(strings.Split(sentences[0], ":"))
	if len(mt) > 1 {
		if len(mt) > 1 {
			return mt[0], mt[1], strings.Join(sentences[1:], " ")
		} else {
			return mt[0], "", strings.Join(sentences[1:], " ")
		}
	} else {
		if len(sentences) > 1 {
			return mt[0], "", sentences[1]
		} else {
			return mt[0], "", ""
		}
	}
}

func arrangeLine(str []string) []string {
	var newList []string
	for _, s := range str {
		newList = append(newList, strings.Trim(s, "\n"))
	}
	return newList
}
