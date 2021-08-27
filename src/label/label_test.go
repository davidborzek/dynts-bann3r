package label

import (
	"os"
	"strings"
	"testing"

	"github.com/multiplay/go-ts3"
	"github.com/stretchr/testify/assert"
)

func TestMain(setup *testing.M) {
	m = map[string]func(*ts3.Client) (string, error){
		"clientsonline": func(c *ts3.Client) (string, error) { return "10", nil },
		"maxclients":    func(c *ts3.Client) (string, error) { return "32", nil },
		"timeHH":        func(c *ts3.Client) (string, error) { return "00", nil },
		"timeMM":        func(c *ts3.Client) (string, error) { return "00", nil },
		"timeSS":        func(c *ts3.Client) (string, error) { return "00", nil },
	}

	marg = map[string]func(*ts3.Client, []string) (string, error){
		"groupcount": func(c *ts3.Client, a []string) (string, error) { return strings.Join(a, ","), nil },
	}

	exitVal := setup.Run()
	os.Exit(exitVal)
}

func TestGenerateLabelReplacesClientsOnlineCorrectly(t *testing.T) {
	replaced, _ := GenerateLabel("%clientsonline%", nil)
	assert.Equal(t, replaced, "10")
}

func TestGenerateLabelReplacesMaxClientsCorrectly(t *testing.T) {
	replaced, _ := GenerateLabel("%maxclients%", nil)
	assert.Equal(t, replaced, "32")
}

func TestGenerateLabelReplacesTimeHHCorrectly(t *testing.T) {
	replaced, _ := GenerateLabel("%timeHH%", nil)
	assert.Equal(t, replaced, "00")
}

func TestGenerateLabelReplacesTimeMMCorrectly(t *testing.T) {
	replaced, _ := GenerateLabel("%timeMM%", nil)
	assert.Equal(t, replaced, "00")
}

func TestGenerateLabelReplacesTimeSSCorrectly(t *testing.T) {
	replaced, _ := GenerateLabel("%timeSS%", nil)
	assert.Equal(t, replaced, "00")
}

func TestGenerateLabelReplacesGroupCountWithArgsCorrectly(t *testing.T) {
	replaced, _ := GenerateLabel("%groupcount[1,2,3,4]%", nil)
	assert.Equal(t, replaced, "1,2,3,4")
}

func TestGenerateLabelReturnsErrorForUnknownPlaceholder(t *testing.T) {
	replaced, err := GenerateLabel("%unknown_placeholer%", nil)
	assert.Equal(t, replaced, "ERROR")
	assert.Equal(t, err.Error(), "placeholder: 'unknown_placeholer' does not exists")
}
