package gpmarker_test

import (
	"testing"

	"gpmarker"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, gpmarker.Analyzer, "a")
}

func TestSplitSentence(t *testing.T) {

}

func TestArrangeLine(t *testing.T) {
	strLine1 := []string{
		"\na\n",
		"b\n",
		"\nc",
	}
	strLine2 := []string{
		"a",
		"b",
		"c\n",
	}
	strLine1Result := gpmarker.ArrangeLine(strLine1)
	strLine2Result := gpmarker.ArrangeLine(strLine2)
	t.Run("改行を消去できているか", func(t *testing.T) {
		fault := false
		for i := range strLine1 {
			if strLine1Result[i] != strLine2Result[i] {
				fault = true
			}
		}
		if fault {
			t.Fatal("改行消去ができていません")
		}
	})
}