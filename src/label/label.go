package label

import (
	"log"
	"regexp"
	"strings"
)

func ReplacePlaceholders(text string, values map[string]string) string {
	placeholderList := getPlaceholderList(text)

	for _, placeholder := range placeholderList {
		p := strings.ReplaceAll(placeholder, "%", "")
		value := values[p]

		if p == "" {
			log.Println("[ERROR] Placeholder '" + placeholder + "' does not exists.")
			return "ERROR"
		}

		text = strings.ReplaceAll(text, placeholder, value)
	}

	return text
}

func getPlaceholderList(text string) []string {
	placeholderRegexp, _ := regexp.Compile("%[^%]+%")
	matches := placeholderRegexp.FindAllString(text, -1)

	return matches
}
