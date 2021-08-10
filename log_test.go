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
			Expected:   "TestLogger_Logf() Hello World!\n",
		},
		{
			Format:     "%s",
			Parameters: []interface{}{""},
			Expected:   "TestLogger_Logf() \n",
		},
		{
			Format:     "",
			Parameters: []interface{}{},
			Expected:   "TestLogger_Logf() \n",
		},
		{
			Format:     "%d + %d = %d",
			Parameters: []interface{}{2, 3, 5},
			Expected:   "TestLogger_Logf() 2 + 3 = 5\n",
		},
		{
			Format:     "My name is %s.",
			Parameters: []interface{}{"Joe"},
			Expected:   "TestLogger_Logf() My name is Joe.\n",
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
			Expected:   "Log() Hello World!\n",
		},
		{
			Parameters: []interface{}{""},
			Expected:   "Log() \n",
		},
		{
			Parameters: []interface{}{},
			Expected:   "Log() \n",
		},
		{
			Parameters: []interface{}{"Hello", " ", "World!"},
			Expected:   "Log() Hello World!\n",
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

func TestLogger_Prefix(t *testing.T) {
	var output strings.Builder
	var logger Logger
	logger.Writer = &output

	prefixedLogger := logger.Prefix("apple", "banana", "cherry")

	prefixedLogger.Log("Hello world with prefixes!")
	expected := "apple: banana: cherry: Log() Hello world with prefixes!\n"
	if output.String() != expected {
		t.Errorf("Log did not log what was expected.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", output.String())
	}
	doublePrefixedLogger := prefixedLogger.Prefix("date")

	doublePrefixedLogger.Log("I am here with more prefixes!")
	expected = "apple: banana: cherry: Log() Hello world with prefixes!\napple: banana: cherry: date: Log() I am here with more prefixes!\n"
	if output.String() != expected {
		t.Errorf("Log did not log what was expected.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", output.String())
	}
}

func TestLogger_Begin(t *testing.T) {
	var output strings.Builder
	var logger Logger
	logger.Writer = &output
	logger.Begin()
	expected := "Log() BEGIN\n"
	if output.String() != expected {
		t.Errorf("Log did not log what was expected.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", output.String())
	}
}

func TestLogger_End(t *testing.T) {
	var output strings.Builder
	var logger Logger
	logger.Writer = &output
	subLogger := logger.Begin()
	subLogger.End()
	expected := "Log() BEGIN\nLog() END δt="
	if !strings.Contains(output.String(), expected) {
		t.Errorf("Log did not log what was expected.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", output.String())
	}
}