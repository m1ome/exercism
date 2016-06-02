package bob

import (
	"strings"
	"unicode"
)

const testVersion = 2

const question, yell, empty, def = "Sure.", "Whoa, chill out!", "Fine. Be that way!", "Whatever."

func Hey(message string) string {
	if len(message) == 0 || len(strings.TrimSpace(message)) == 0 {
		return empty
	}

	for _, r := range message {
		if unicode.IsLetter(r) {
			if strings.ToUpper(message) == message {
				return yell
			} else {
				break
			}
		}
	}

	if message[len(message)-1] == '?' {
		return question
	}

	return def
}
