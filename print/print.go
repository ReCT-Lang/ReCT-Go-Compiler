package print

import (
	"fmt"
)

// this is just a little helper for colored printing

const (
	Black      = "\033[1;30m%s\033[0m" // &bl
	DarkRed    = "\033[1;31m%s\033[0m" // &dr
	DarkGreen  = "\033[1;32m%s\033[0m" // &dg
	DarkYellow = "\033[1;33m%s\033[0m" // &dy
	DarkBlue   = "\033[1;34m%s\033[0m" // &db
	Magenta    = "\033[1;35m%s\033[0m" // &m
	DarkCyan   = "\033[1;36m%s\033[0m" // &dc
	Gray       = "\033[1;37m%s\033[0m" // &g
	DarkGray   = "\033[1;90m%s\033[0m" // &dgr
	Red        = "\033[1;91m%s\033[0m" // &r
	Green      = "\033[1;92m%s\033[0m" // &g
	Yellow     = "\033[1;93m%s\033[0m" // &y
	Blue       = "\033[1;94m%s\033[0m" // &b
	Pink       = "\033[1;95m%s\033[0m" // &p
	Cyan       = "\033[1;96m%s\033[0m" // &c
	White      = "\033[1;97m%s\033[0m" // &w

	EBlack      = "\033[1;30m" // All E variants are "color + &e"
	EDarkRed    = "\033[1;31m" // like &e&bl for EBlack
	EDarkGreen  = "\033[1;32m" // for EReset use &%
	EDarkYellow = "\033[1;33m"
	EDarkBlue   = "\033[1;34m"
	EMagenta    = "\033[1;35m"
	EDarkCyan   = "\033[1;36m"
	EGray       = "\033[1;37m"
	EDarkGray   = "\033[1;90m"
	ERed        = "\033[1;91m"
	EGreen      = "\033[1;92m"
	EYellow     = "\033[1;93m"
	EBlue       = "\033[1;94m"
	EPink       = "\033[1;95m"
	ECyan       = "\033[1;96m"
	EWhite      = "\033[1;97m"
	EReset      = "\033[0m"
)

func PrintC(color string, message string) {
	fmt.Printf(color, message+"\n")
}

func PrintCF(color string, message string, slotins ...interface{}) {
	text := fmt.Sprintf(message, slotins...)
	fmt.Printf(color, text+"\n")
}

func WriteC(color string, message string) {
	fmt.Printf(color, message)
}

func WriteCF(color string, message string, slotins ...interface{}) {
	text := fmt.Sprintf(message, slotins...)
	fmt.Printf(color, text)
}
