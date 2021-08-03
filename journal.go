package gpmarker

import "fmt"

type (
	Journal struct {
		Type string
		FileName string
		Pos int
		Text string
		CodeText []string
	}
	JournalList struct {
		InfoList []*Journal
		WarnList []*Journal
		OtherList []*Journal
	}
	JournalType string
)

const (
	Info JournalType = "info"
	Warning JournalType = "warn"
)

func (jl *JournalList) Add(j *Journal) {
	switch JournalType(j.Type) {
	case Info:
		jl.InfoList = append(jl.InfoList, j)
	case Warning:
		jl.WarnList = append(jl.WarnList, j)
	default:
		jl.OtherList = append(jl.OtherList, j)
	}
}

func (jl *JournalList) Preview(requireRows bool) {
	if len(jl.InfoList) == 0 && len(jl.WarnList) == 0 && len(jl.OtherList) == 0 {
		PrintAny(PGreen, "No Exist marker")
		return
	}
	fmt.Println("")
	format := "[%s.go:%d] => %s\n"
	formatWithoutText := "[%s.go:%d]\n"
	for i, w := range jl.WarnList {
		if i == 0 { PrintAny(PMagenta, "=== Warnings") }
		if w.Text == "" {
			fmt.Printf(formatWithoutText, w.FileName, w.Pos)
			if requireRows { printRows(w.CodeText, PYellow, PGreen) }
		} else {
			fmt.Printf(format, w.FileName, w.Pos, w.Text)
			if requireRows { printRows(w.CodeText, PYellow, PGreen) }
		}
		if i == len(jl.WarnList) - 1 { fmt.Println("") }
	}

	for i, w := range jl.InfoList {
		if i == 0 { PrintAny(PCyan, "=== Information") }
		if w.Text == "" {
			fmt.Printf(formatWithoutText, w.FileName, w.Pos)
			if requireRows { printRows(w.CodeText, PYellow, PGreen) }
		} else {
			fmt.Printf(format, w.FileName, w.Pos, w.Text)
			if requireRows { printRows(w.CodeText, PYellow, PGreen) }
		}
		if i == len(jl.InfoList) - 1 { fmt.Println("") }
	}

	for i, w := range jl.OtherList {
		if i == 0 { PrintAny(PBlue, "=== Others") }
		if w.Text == "" {
			fmt.Printf(formatWithoutText, w.FileName, w.Pos)
			if requireRows { printRows(w.CodeText, PYellow, PGreen) }
		} else {
			fmt.Printf(format, w.FileName, w.Pos, w.Text)
			if requireRows { printRows(w.CodeText, PYellow, PGreen) }
		}
		if i == len(jl.OtherList) - 1 { fmt.Println("") }
	}
}

// private

func printRows(rows []string, colors ...int) {
	preColor := PYellow
	for i, row := range rows {
		var colorNum int
		if len(colors) > i {
			colorNum, preColor = colors[i], colors[i]
		} else {
			colorNum = preColor
		}
		PrintAny(colorNum, row)
	}
}
