package gpmarker

import "fmt"

const (
	PBlack = iota + 30
	PRed
	PGreen
	PYellow
	PBlue
	PMagenta
	PCyan
	PWhite
)

func PrintAny(number int, str string) {
	fmt.Printf("\x1b[%dm%s\x1b[0m\n", number, str)
}
