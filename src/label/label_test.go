package label_test

import (
	"dynts-bann3r/src/label"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPlaceholderValueMap = map[string]string{
	"MAX_CLIENTS":          "32",
	"REAL_CLIENTS_ONLINE":  "7",
	"CLIENTS_ONLINE":       "8",
	"QUERY_CLIENTS_ONLINE": "1",
	"ADMIN_CLIENTS_ONLINE": "3",
	"SERVER_NAME":          "A server",
	"SERVER_PORT":          "9987",
	"TIME_HH":              "12",
	"TIME_MM":              "45",
	"TIME_SS":              "13",
}

func Test_ReplacePlaceholders_Replaces_MAX_CLIENTS_Correctly(t *testing.T) {
	replaced := label.ReplacePlaceholders("%MAX_CLIENTS%", testPlaceholderValueMap)
	assert.Equal(t, replaced, "32")
}

func Test_ReplacePlaceholders_Replaces_REAL_ONLINE_CLIENTS_Correctly(t *testing.T) {
	replaced := label.ReplacePlaceholders("%REAL_CLIENTS_ONLINE%", testPlaceholderValueMap)
	assert.Equal(t, replaced, "7")
}

func Test_ReplacePlaceholders_Replaces_CLIENTS_ONLINE_Correctly(t *testing.T) {
	replaced := label.ReplacePlaceholders("%CLIENTS_ONLINE%", testPlaceholderValueMap)
	assert.Equal(t, replaced, "8")
}

func Test_ReplacePlaceholders_Replaces_QUERY_CLIENTS_ONLINE_Correctly(t *testing.T) {
	replaced := label.ReplacePlaceholders("%QUERY_CLIENTS_ONLINE%", testPlaceholderValueMap)
	assert.Equal(t, replaced, "1")
}

func Test_ReplacePlaceholders_Replaces_ADMIN_CLIENTS_ONLINE_Correctly(t *testing.T) {
	replaced := label.ReplacePlaceholders("%ADMIN_CLIENTS_ONLINE%", testPlaceholderValueMap)
	assert.Equal(t, replaced, "3")
}

func Test_ReplacePlaceholders_Replaces_SERVER_NAME_Correctly(t *testing.T) {
	replaced := label.ReplacePlaceholders("%SERVER_NAME%", testPlaceholderValueMap)
	assert.Equal(t, replaced, "A server")
}

func Test_ReplacePlaceholders_Replaces_SERVER_PORT_Correctly(t *testing.T) {
	replaced := label.ReplacePlaceholders("%SERVER_PORT%", testPlaceholderValueMap)
	assert.Equal(t, replaced, "9987")
}

func Test_ReplacePlaceholders_Replaces_TIME_HH_Correctly(t *testing.T) {
	replaced := label.ReplacePlaceholders("%TIME_HH%", testPlaceholderValueMap)
	assert.Equal(t, replaced, "12")
}

func Test_ReplacePlaceholders_Replaces_TIME_MM_Correctly(t *testing.T) {
	replaced := label.ReplacePlaceholders("%TIME_MM%", testPlaceholderValueMap)
	assert.Equal(t, replaced, "45")
}

func Test_ReplacePlaceholders_Replaces_TIME_SS_Correctly(t *testing.T) {
	replaced := label.ReplacePlaceholders("%TIME_SS%", testPlaceholderValueMap)
	assert.Equal(t, replaced, "13")
}
