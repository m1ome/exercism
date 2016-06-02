package bob

import (
	"strings"
	"unicode"
)

const testVersion = 2

var question, yell, empty, def = "Sure.", "Whoa, chill out!", "Fine. Be that way!", "Whatever."

func Hey(message string) string {
	if strings.TrimSpace(message) == "" {
		return empty
	}

	var let, upcase = false, true
	for _, r := range message {
		if unicode.IsLetter(r) {
			let = true
			if !unicode.IsUpper(r) {
				upcase = false
				break
			}
		}
	}

	if let && upcase {
		return yell
	}

	if strings.HasSuffix(message, "?") {
		return question
	}

	return def
}
