package utils

import "strings"

func ToCamelCase(input string) string {
	//keep everything in lower case
	input = strings.ToLower(input)

	// Replace all non-alphanumeric characters with spaces
	space := func(r rune) rune {
		if !('a' <= r && r <= 'z') && !('A' <= r && r <= 'Z') && !('0' <= r && r <= '9') {
			return ' '
		}
		return r
	}
	s := strings.Map(space, input)

	// Split the modified string into individual words
	words := strings.Fields(s)

	// Capitalize the first letter of each word (except the first word)
	for i := 0; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}

	// Join the words to form the camel case string
	camelCase := strings.Join(words, " ")

	return camelCase
}
