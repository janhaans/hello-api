package translation

import "strings"

func Translate(word string, language string) string {
	word = SanitizeInput(word)
	language = SanitizeInput(language)
	if word != "hello" {
		return ""
	}

	switch language {
	case "english":
		return "hello"
	case "finnish":
		return "hei"
	case "german":
		return "hallo"
	case "french":
		return "bonjour"
	default:
		return ""
	}
}

func SanitizeInput(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	return s
}
