package print

import "fmt"

// this is just a little helper for colored printing

const (
	Black      = "\033[1;30m%s\033[0m"
	DarkRed    = "\033[1;31m%s\033[0m"
	DarkGreen  = "\033[1;32m%s\033[0m"
	DarkYellow = "\033[1;33m%s\033[0m"
	DarkBlue   = "\033[1;34m%s\033[0m"
	Magenta    = "\033[1;35m%s\033[0m"
	DarkCyan   = "\033[1;36m%s\033[0m"
	Gray       = "\033[1;37m%s\033[0m"
	DarkGray   = "\033[1;90m%s\033[0m"
	Red        = "\033[1;91m%s\033[0m"
	Green      = "\033[1;92m%s\033[0m"
	Yellow     = "\033[1;93m%s\033[0m"
	Blue       = "\033[1;94m%s\033[0m"
	Pink       = "\033[1;95m%s\033[0m"
	Cyan       = "\033[1;96m%s\033[0m"
	White      = "\033[1;97m%s\033[0m"

	EBlack      = "\033[1;30m"
	EDarkRed    = "\033[1;31m"
	EDarkGreen  = "\033[1;32m"
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
