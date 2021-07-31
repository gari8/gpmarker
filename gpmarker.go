package gpmarker

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

const doc = "gpmarker is ..."
var journalList JournalList

// SplitSentence mark, type, text
func SplitSentence(s string) (string, string, string) {
	sentences := ArrangeLine(strings.Split(s, " "))
	mt := ArrangeLine(strings.Split(sentences[0], ":"))
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

func ArrangeLine(str []string) []string {
	var newList []string
	for _, s := range str {
		newList = append(newList, strings.Trim(s, "\n"))
	}
	return newList
}

// Walk用の実行関数
func exec(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() && filepath.Ext(path) == ".go" {
		fSet := token.NewFileSet()
		file, err := parser.ParseFile(fSet, path, nil, parser.ParseComments)
		if err != nil {
			return err
		}
		cMap := ast.NewCommentMap(fSet, file, file.Comments)
		for _, c := range cMap.Comments() {
			mark, tp, text := SplitSentence(c.Text())
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
	return nil
}

func WalkDirectory(fp string) (*JournalList, error) {
	if err := filepath.Walk(fp, exec); err != nil {
		PrintAny(PRed, "error : file path not found")
		return nil, err
	}
	return &journalList, nil
}
