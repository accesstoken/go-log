package log

import (
	"strings"
	"testing"
)

func TestLogger_Logf(t *testing.T) {

	tests := []struct {
		Format     string
		Parameters []interface{}
		Expected   string
	}{
		{
			Format:     "%s",
			Parameters: []interface{}{"Hello World!"},
			Expected:   "Hello World!\n",
		},
		{
			Format:     "%s",
			Parameters: []interface{}{""},
			Expected:   "\n",
		},
		{
			Format:     "",
			Parameters: []interface{}{},
			Expected:   "\n",
		},
		{
			Format:     "%d + %d = %d",
			Parameters: []interface{}{2, 3, 5},
			Expected:   "2 + 3 = 5\n",
		},
		{
			Format:     "My name is %s.",
			Parameters: []interface{}{"Joe"},
			Expected:   "My name is Joe.\n",
		},
	}

	for testNumber, test := range tests {

		var output strings.Builder
		var logger Logger
		logger.Writer = &output

		logger.Logf(test.Format, test.Parameters...)

		if expected, actual := test.Expected, output.String(); expected != actual {
			t.Errorf("For test #%d, Logf did not log what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

	}
}

func TestLogger_Log(t *testing.T) {

	tests := []struct {
		Parameters []interface{}
		Expected   string
	}{
		{
			Parameters: []interface{}{"Hello World!"},
			Expected:   "Hello World!\n",
		},
		{
			Parameters: []interface{}{""},
			Expected:   "\n",
		},
		{
			Parameters: []interface{}{},
			Expected:   "\n",
		},
		{
			Parameters: []interface{}{"Hello", " ", "World!"},
			Expected:   "Hello World!\n",
		},
	}

	for testNumber, test := range tests {

		var output strings.Builder
		var logger Logger
		logger.Writer = &output

		logger.Log(test.Parameters...)

		if expected, actual := test.Expected, output.String(); expected != actual {
			t.Errorf("For test #%d, Log did not log what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

	}
}