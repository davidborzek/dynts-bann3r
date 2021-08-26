package label

import (
	"dynts-bann3r/src/teamspeak"
	"errors"
	"regexp"
	"strings"

	"github.com/multiplay/go-ts3"
)

var m = map[string]func(*ts3.Client) (string, error){
	"clientsonline": teamspeak.CountOnlineClients,
	"maxclients":    teamspeak.GetMaxClients,
}

var marg = map[string]func(*ts3.Client, []string) (string, error){
	"groupcount": teamspeak.CountOnlineClientsInGroups,
}

func replacePlaceholder(placeholder string, client *ts3.Client) (string, error) {
	fn := m[placeholder]

	if fn == nil {
		return "", errors.New("placeholder: " + placeholder + "does not exists")
	}

	return fn(client)
}

func replacePlaceholderWithArguments(placeholder string, client *ts3.Client, args []string) (string, error) {
	fn := marg[placeholder]

	if fn == nil {
		return "", errors.New("placeholder: " + placeholder + "does not exists")
	}

	return fn(client, args)
}

func GenerateLabel(text string, client *ts3.Client) (string, error) {
	r, _ := regexp.Compile("%[^%]+%")
	z, _ := regexp.Compile(`\[[^\[\]]+\]`)

	a := r.FindAllString(text, -1)

	for _, placeholder := range a {
		p := placeholder

		b := z.FindAllString(p, 1)

		var params []string

		if len(b) > 0 {
			paramsStr := strings.ReplaceAll(b[0], "[", "")
			paramsStr = strings.ReplaceAll(paramsStr, "]", "")

			p = strings.ReplaceAll(p, b[0], "")

			params = strings.Split(paramsStr, ",")
		}

		var replaced string
		var err error

		if len(params) > 0 {
			replaced, err = replacePlaceholderWithArguments(strings.ReplaceAll(p, "%", ""), client, params)
		} else {
			replaced, err = replacePlaceholder(strings.ReplaceAll(p, "%", ""), client)
		}

		if err != nil {
			return "ERROR", err
		}

		text = strings.ReplaceAll(text, placeholder, replaced)
	}

	return text, nil
}
