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
	Gray       = "\033[1;37m%s\033[0m" // &gr
	DarkGray   = "\033[1;90m%s\033[0m" // &dgr
	Red        = "\033[1;91m%s\033[0m" // &r
	Green      = "\033[1;92m%s\033[0m" // &g
	Yellow     = "\033[1;93m%s\033[0m" // &y
	Blue       = "\033[1;94m%s\033[0m" // &b
	Pink       = "\033[1;95m%s\033[0m" // &p
	Cyan       = "\033[1;96m%s\033[0m" // &c
	White      = "\033[1;97m%s\033[0m" // &w
	// NOTE: if you want to use write '&' use "&&"

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

func WriteCF(color string, message string, slotins ...interface{}) {
	text := fmt.Sprintf(message, slotins...)
	fmt.Printf(color, text)
}

// Format multiline coloring made easier!
// use colour codes in the string like "&pHello gamers, %rHow are you doing today&dgr?"
// and Format will colour code it for you.
// the defaultColour will be the default colour of text if it is not coloured
// also doing the same colour code will reset the colour e.g. "Hello, &rW&rorld!" only "W" will be red
func Format(message string, defaultColour string, a ...interface{}) string {

	var buffer string                 // Final buffer which is returned
	var temp string                   // This stores the current coloured text like "Hello gamers, "
	var currentColour = defaultColour // Stores the color, this changes depending on the code.
	skip := 0                         // Sometimes we need to skip the next iteration (can be more than one)

	for index, letter := range []rune(message) { // yes, I'm using bytes, if you want runes just change to a []rune() cast

		// This is because we have a letter after the '&' we are processing, so we have to skip it in the next iteration
		// in cases where we have more than one letter we'll be skipping multiple iterations (hence why skip is an integer).
		if skip > 0 {
			skip--
			continue

		} else if letter == '&' { // This is where we process the colour code
			index++
			if index < len(message) {

				skip++ // We know it must be a colour code or "&&"
				var newColor string

				// Most of this is self-explanatory, though it is a bit hard to read.
				switch message[index] {
				case 'r':
					newColor = Red
				case 'g':
					// Handling multi-character color codes
					if index++; index < len(message) && message[index] == 'r' {
						newColor = Gray
						skip++
					} else {
						newColor = Green
					}
				case 'y':
					newColor = Yellow
				case 'b':
					// Handling multi-character color codes
					if index++; index < len(message) && message[index] == 'l' {
						newColor = Black
						skip++
					} else {
						newColor = Blue
					}
				case 'c':
					newColor = Cyan
				case 'w':
					newColor = White
				case 'm':
					newColor = Magenta
				case 'p':
					newColor = Pink
				case 'd':
					index++
					skip++
					// Handling multi-character color codes
					switch message[index] {
					case 'r':
						newColor = DarkRed
					case 'y':
						newColor = DarkYellow
					case 'b':
						newColor = DarkBlue
					case 'g':
						// Handling multi-character color codes
						if index++; index < len(message) && message[index] == 'r' {
							newColor = DarkGray
							skip++
						} else {
							newColor = DarkGreen
						}
					case 'c':
						newColor = DarkCyan
					default:
						newColor = defaultColour // If we get a character we are not expecting we just go to default
					}
				case '&':
					newColor = currentColour
					temp += "&"
				default:
					newColor = defaultColour
				}

				// Append temp to buffer (this was the previous color)
				buffer += fmt.Sprintf(currentColour, temp)

				// Set up new colour
				if currentColour != newColor {
					currentColour = newColor
				} else {
					currentColour = defaultColour // If we get a character we are not expecting we just go to default
				}
				temp = "" // Reset temp for new content of course
			}

		} else { // If it isn't a colour code, we just append the letter onto the temporary buffer
			temp += string(rune(letter))
		}
	}
	// We need to append the final temp buffer
	buffer += fmt.Sprintf(currentColour, temp)
	return fmt.Sprintf(buffer, a...) // Epic we all done (also format any values)
}
