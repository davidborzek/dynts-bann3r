package label

import (
	"os"
	"strings"
	"testing"

	"github.com/multiplay/go-ts3"
	"github.com/stretchr/testify/assert"
)

func TestMain(setup *testing.M) {
	tsPlaceholderFunctionMap = map[string]func(*ts3.Client) (string, error){
		"clientsonline": func(c *ts3.Client) (string, error) { return "10", nil },
		"maxclients":    func(c *ts3.Client) (string, error) { return "32", nil },
		"servername":    func(c *ts3.Client) (string, error) { return "dynts-bann3r", nil },
		"serverport":    func(c *ts3.Client) (string, error) { return "9987", nil },
	}

	tsPlaceholderArgumentFunctionMap = map[string]func(*ts3.Client, []string) (string, error){
		"groupcount": func(c *ts3.Client, a []string) (string, error) { return strings.Join(a, ","), nil },
	}

	timePlaceholderFunctionMap = map[string]func() string{
		"timeHH": func() string { return "00" },
		"timeMM": func() string { return "00" },
		"timeSS": func() string { return "00" },
	}

	exitVal := setup.Run()
	os.Exit(exitVal)
}

func TestGenerateLabelReplacesClientsOnlineCorrectly(t *testing.T) {
	replaced := GenerateLabel("%clientsonline%", nil)
	assert.Equal(t, replaced, "10")
}

func TestGenerateLabelReplacesMaxClientsCorrectly(t *testing.T) {
	replaced := GenerateLabel("%maxclients%", nil)
	assert.Equal(t, replaced, "32")
}

func TestGenerateLabelReplacesServernameCorrectly(t *testing.T) {
	replaced := GenerateLabel("%servername%", nil)
	assert.Equal(t, replaced, "dynts-bann3r")
}

func TestGenerateLabelReplacesServerportCorrectly(t *testing.T) {
	replaced := GenerateLabel("%serverport%", nil)
	assert.Equal(t, replaced, "9987")
}

func TestGenerateLabelReplacesTimeHHCorrectly(t *testing.T) {
	replaced := GenerateLabel("%timeHH%", nil)
	assert.Equal(t, replaced, "00")
}

func TestGenerateLabelReplacesTimeMMCorrectly(t *testing.T) {
	replaced := GenerateLabel("%timeMM%", nil)
	assert.Equal(t, replaced, "00")
}

func TestGenerateLabelReplacesTimeSSCorrectly(t *testing.T) {
	replaced := GenerateLabel("%timeSS%", nil)
	assert.Equal(t, replaced, "00")
}

func TestGenerateLabelReplacesGroupCountWithArgsCorrectly(t *testing.T) {
	replaced := GenerateLabel("%groupcount[1,2,3,4]%", nil)
	assert.Equal(t, replaced, "1,2,3,4")
}

func TestGenerateLabelReturnsErrorForUnknownPlaceholder(t *testing.T) {
	replaced := GenerateLabel("%unknown_placeholer%", nil)
	assert.Equal(t, replaced, "ERROR")
}
