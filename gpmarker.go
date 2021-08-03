package gpmarker

import (
	"bufio"
	"fmt"
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
			return mt[0], "", strings.Join(sentences[1:], " ")
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
	if info.IsDir() || filepath.Ext(path) != ".go" {
		return nil
	}
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	rows := getContentFromFile(file)

	fSet := token.NewFileSet()
	pf, err := parser.ParseFile(fSet, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	cMap := ast.NewCommentMap(fSet, pf, pf.Comments)
	for _, c := range cMap.Comments() {
		ff := fSet.File(c.Pos())
		mark, tp, text := SplitSentence(c.Text())
		line := ff.Line(c.Pos())
		endLine := ff.Line(c.End())
		if mark == "mark" {
			journal := Journal{
				Type: tp,
				FileName: pf.Name.String(),
				Pos: line,
				Text: text,
				CodeText: rows[line-1:endLine+2],
			}
			journalList.Add(&journal)
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

// private

func getContentFromFile(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var result []string
	i := 1
	for scanner.Scan() {
		s := fmt.Sprintf("%d: %s", i, scanner.Text())
		result = append(result, s)
		i++
	}
	return result
}
