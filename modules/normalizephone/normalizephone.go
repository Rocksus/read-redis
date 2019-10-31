package normalizephone

import "fmt"

// Normalize number to format 62xxx
func Normalize(number string) string {
	// Check if its a valid phone number
	if len(number) < 8 {
		return "-"
	}
	if number[0] == '+' {
		number = number[1:len(number)]
	}
	if number[0:2] == "62" {
		return number
	} else if number[0:2] == "08" {
		number = fmt.Sprintf("%s%s", "62", number[1:len(number)])
		return number
	}
	return "-"
}
