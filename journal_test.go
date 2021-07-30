package gpmarker

import "testing"

var testCaseA = &Journal{
	Type: "",
	FileName: "a",
	Pos: 30213,
	Text: "abc abc",
}

var testCaseB = &Journal{
	Type: "info",
	FileName: "b",
	Pos: 30213,
	Text: "abc abc",
}

func TestJournalList_Add(t *testing.T) {
	var jl JournalList
	jl.Add(testCaseA)
	jl.Add(testCaseB)
	t.Run("タイプ毎の正しい仕訳", func(t *testing.T) {
		if !(len(jl.InfoList) == 1 && len(jl.OtherList) == 1 && len(jl.WarnList) == 0) {
			t.Fatal("仕訳が正しくできていません")
		}
	})
}
