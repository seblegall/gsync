package prompt

import (
	tm "github.com/buger/goterm"
)

func Title(text string) {
	tm.Println()
	tm.Println(tm.Color(tm.Bold(text),tm.GREEN))
	tm.Println()
	tm.Flush()
}

func Warning(text string) {
	tm.Println()
	tm.Println(tm.Color(tm.Bold(text),tm.YELLOW))
	tm.Println()
	tm.Flush()
}

func Info(text string) {
	tm.Println(tm.Color(text,tm.WHITE))
	tm.Flush()
}

func UserError(text string) {
	tm.Println(tm.Color(tm.Bold(text),tm.RED))
	tm.Flush()
}


