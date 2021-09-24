package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {

	word = strings.ToLower(word)

	for i := range word {
		if strings.Count(word, string(word[i])) > 1 && string(word[i]) != "-" && string(word[i]) != " " {
			return false
		}
	}
	return true
}
