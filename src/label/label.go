package label

import (
	"dynts-bann3r/src/teamspeak"
	"dynts-bann3r/src/utils"
	"errors"
	"log"
	"regexp"
	"strings"

	"github.com/multiplay/go-ts3"
)

var tsPlaceholderFunctionMap = map[string]func(*ts3.Client) (string, error){
	"clientsonline": teamspeak.CountOnlineClients,
	"maxclients":    teamspeak.GetMaxClients,
}

var tsPlaceholderArgumentFunctionMap = map[string]func(*ts3.Client, []string) (string, error){
	"groupcount": teamspeak.CountOnlineClientsInGroups,
}

var timePlaceholderFunctionMap = map[string]func() string{
	"timeHH": utils.GetHours,
	"timeMM": utils.GetMinutes,
	"timeSS": utils.GetSeconds,
}

var replacePlaceholder = func(placeholder string, client *ts3.Client, args []string) (string, error) {
	placeholder = strings.ReplaceAll(placeholder, "%", "")

	if len(args) > 0 {
		placeholder = strings.ReplaceAll(placeholder, "["+strings.Join(args, ",")+"]", "")

		fn := tsPlaceholderArgumentFunctionMap[placeholder]

		if fn == nil {
			return "", errors.New("placeholder with arguments: '" + placeholder + "' does not exists")
		}

		return fn(client, args)
	}

	if fn := tsPlaceholderFunctionMap[placeholder]; fn == nil {
		timeFn := timePlaceholderFunctionMap[placeholder]

		if timeFn == nil {
			return "", errors.New("placeholder: '" + placeholder + "' does not exists")
		}

		return timeFn(), nil

	} else {
		return fn(client)

	}
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

func GenerateLabel(text string, client *ts3.Client) string {
	placeholderList := getPlaceholderList(text)

	for _, placeholder := range placeholderList {
		args := getArguments(placeholder)

		replaced, err := replacePlaceholder(placeholder, client, args)

		if err != nil {
			log.Println(err)
			return "ERROR"
		}

		text = strings.ReplaceAll(text, placeholder, replaced)
	}

	return text
}
