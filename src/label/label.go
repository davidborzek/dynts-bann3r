package label

import (
	"dynts-bann3r/src/teamspeak"
	"dynts-bann3r/src/utils"
	"errors"
	"regexp"
	"strings"

	"github.com/multiplay/go-ts3"
)

var m = map[string]func(*ts3.Client) (string, error){
	"clientsonline": teamspeak.CountOnlineClients,
	"maxclients":    teamspeak.GetMaxClients,
	"timeHH":        utils.GetHours,
	"timeMM":        utils.GetMinutes,
	"timeSS":        utils.GetSeconds,
}

var marg = map[string]func(*ts3.Client, []string) (string, error){
	"groupcount": teamspeak.CountOnlineClientsInGroups,
}

var replacePlaceholder = func(placeholder string, client *ts3.Client) (string, error) {
	fn := m[placeholder]

	if fn == nil {
		return "", errors.New("placeholder: '" + placeholder + "' does not exists")
	}

	return fn(client)
}

func replacePlaceholderWithArguments(placeholder string, client *ts3.Client, args []string) (string, error) {
	fn := marg[placeholder]

	if fn == nil {
		return "", errors.New("placeholder: '" + placeholder + "' does not exists")
	}

	return fn(client, args)
}

func getPlaceholderList(text string) []string {
	placeholderRegexp, _ := regexp.Compile("%[^%]+%")
	matches := placeholderRegexp.FindAllString(text, -1)

	return matches
}

func getArguments(placeholder string) []string {
	argumentRegexp, _ := regexp.Compile(`\[[^\[\]]+\]`)
	matches := argumentRegexp.FindAllString(placeholder, 1)

	var args []string

	if len(matches) > 0 {
		argStr := strings.ReplaceAll(matches[0], "[", "")
		argStr = strings.ReplaceAll(argStr, "]", "")

		args = strings.Split(argStr, ",")
	}

	return args
}

func GenerateLabel(text string, client *ts3.Client) (string, error) {
	placeholderList := getPlaceholderList(text)

	for _, placeholder := range placeholderList {
		p := placeholder

		args := getArguments(p)

		var replaced string
		var err error

		if len(args) > 0 {
			p := strings.ReplaceAll(p, "["+strings.Join(args, ",")+"]", "")
			replaced, err = replacePlaceholderWithArguments(strings.ReplaceAll(p, "%", ""), client, args)
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
