package gpmarker

import "fmt"

type (
	Journal struct {
		Type string
		FileName string
		Pos int
		Text string
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

func (jl *JournalList) Preview() {
	fmt.Println("")
	format := "l.%d [%s.go]: %s\n"
	formatWithoutText := "l.%d [%s.go]\n"
	for i, w := range jl.WarnList {
		if i == 0 { PrintAny(PMagenta, "=== Warnings") }
		if w.Text == "" {
			fmt.Printf(formatWithoutText, w.Pos, w.FileName)
		} else {
			fmt.Printf(format, w.Pos, w.FileName, w.Text)
		}
		if i == len(jl.WarnList) - 1 { fmt.Println("") }
	}

	for i, w := range jl.InfoList {
		if i == 0 { PrintAny(PCyan, "=== Information") }
		if w.Text == "" {
			fmt.Printf(formatWithoutText, w.Pos, w.FileName)
		} else {
			fmt.Printf(format, w.Pos, w.FileName, w.Text)
		}
		if i == len(jl.InfoList) - 1 { fmt.Println("") }
	}

	for i, w := range jl.OtherList {
		if i == 0 { PrintAny(PBlue, "=== Others") }
		if w.Text == "" {
			fmt.Printf(formatWithoutText, w.Pos, w.FileName)
		} else {
			fmt.Printf(format, w.Pos, w.FileName, w.Text)
		}
		if i == len(jl.OtherList) - 1 { fmt.Println("") }
	}
}
