package iteration

import "strings"

const repeatCount = 5

// Repeat returns a string consisting of the input character repeated
func Repeat(character string, numberOfRepeats int) string {
	var repeated strings.Builder
	for i := 0; i < numberOfRepeats; i++ {
		repeated.WriteString(character)
	}

	return repeated.String()
}
