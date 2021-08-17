package log

import (
	"strings"
	"testing"
)

func TestLogger_writef(t *testing.T) {

	tests := []struct {
		Format     string
		Parameters []interface{}
		Expected   string
	}{
		{
			Format:     "%s",
			Parameters: []interface{}{"Hello World!"},
			Expected:   "TestLogger_writef() -> Hello World!\n",
		},
		{
			Format:     "%s",
			Parameters: []interface{}{""},
			Expected:   "TestLogger_writef() -> \n",
		},
		{
			Format:     "",
			Parameters: []interface{}{},
			Expected:   "TestLogger_writef() -> \n",
		},
		{
			Format:     "%d + %d = %d",
			Parameters: []interface{}{2, 3, 5},
			Expected:   "TestLogger_writef() -> 2 + 3 = 5\n",
		},
		{
			Format:     "My name is %s.",
			Parameters: []interface{}{"Joe"},
			Expected:   "TestLogger_writef() -> My name is Joe.\n",
		},
	}

	for testNumber, test := range tests {

		var output strings.Builder
		var logger Logger
		logger.Writer = &output

		logger.writef(test.Format, test.Parameters...)

		if expected, actual := test.Expected, output.String(); expected != actual {
			t.Errorf("For test #%d, writef did not log what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

	}
}

func TestLogger_write(t *testing.T) {

	tests := []struct {
		Parameters []interface{}
		Expected   string
	}{
		{
			Parameters: []interface{}{"Hello World!"},
			Expected:   "TestLogger_write() -> Hello World!\n",
		},
		{
			Parameters: []interface{}{""},
			Expected:   "TestLogger_write() -> \n",
		},
		{
			Parameters: []interface{}{},
			Expected:   "TestLogger_write() -> \n",
		},
		{
			Parameters: []interface{}{"Hello", " ", "World!"},
			Expected:   "TestLogger_write() -> Hello World!\n",
		},
	}

	for testNumber, test := range tests {

		var output strings.Builder
		var logger Logger
		logger.Writer = &output

		logger.write(test.Parameters...)

		if expected, actual := test.Expected, output.String(); expected != actual {
			t.Errorf("For test #%d, write did not log what was expected.", testNumber)
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

	prefixedLogger.write("Hello world with prefixes!")

	expected := "TestLogger_Prefix() -> apple: banana: cherry: Hello world with prefixes!\n"
	if output.String() != expected {
		t.Errorf("write did not log what was expected.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", output.String())
		return
	}

	doublePrefixedLogger := prefixedLogger.Prefix("date")

	doublePrefixedLogger.write("I am here with more prefixes!")

	expected = expected + "TestLogger_Prefix() -> apple: banana: cherry: date: I am here with more prefixes!\n"
	if output.String() != expected {
		t.Errorf("write did not log what was expected.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", output.String())
	}
}

func TestLogger_Begin(t *testing.T) {
	var output strings.Builder
	var logger Logger
	logger.Writer = &output
	logger.Begin()
	expected := "TestLogger_Begin() -> BEGIN\n"
	if output.String() != expected {
		t.Errorf("write did not log what was expected.")
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
	expected := "TestLogger_End() -> BEGIN\nTestLogger_End() -> END δt="
	if !strings.Contains(output.String(), expected) {
		t.Errorf("write did not log what was expected.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", output.String())
	}
}

func TestLogger_Alert(t *testing.T) {

	tests := []struct {
		level      uint8
		Parameters []interface{}
		Expected   string
	}{
		{
			level:      0,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      1,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "TestLogger_Alert() -> Hello World!\n",
		},
	}

	for testNumber, test := range tests {

		var output strings.Builder
		var logger Logger
		logger.Writer = &output
		logger = logger.Level(test.level)
		logger.Alert(test.Parameters...)

		if expected, actual := test.Expected, output.String(); expected != actual {
			t.Errorf("For test #%d, Alert did not log what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

	}
}


func TestLogger_Warn(t *testing.T) {

	tests := []struct {
		level      uint8
		Parameters []interface{}
		Expected   string
	}{
		{
			level:      0,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      1,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      2,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "TestLogger_Warn() -> Hello World!\n",
		},
	}

	for testNumber, test := range tests {

		var output strings.Builder
		var logger Logger
		logger.Writer = &output
		logger = logger.Level(test.level)
		logger.Warn(test.Parameters...)

		if expected, actual := test.Expected, output.String(); expected != actual {
			t.Errorf("For test #%d, Warn did not log what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

	}
}


func TestLogger_Highlight(t *testing.T) {

	tests := []struct {
		level      uint8
		Parameters []interface{}
		Expected   string
	}{
		{
			level:      0,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      1,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      2,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      3,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "TestLogger_Highlight() -> Hello World!\n",
		},
	}

	for testNumber, test := range tests {

		var output strings.Builder
		var logger Logger
		logger.Writer = &output
		logger = logger.Level(test.level)
		logger.Highlight(test.Parameters...)

		if expected, actual := test.Expected, output.String(); expected != actual {
			t.Errorf("For test #%d, Highlight did not log what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

	}
}

func TestLogger_Inform(t *testing.T) {

	tests := []struct {
		level      uint8
		Parameters []interface{}
		Expected   string
	}{
		{
			level:      0,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      1,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      2,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      3,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      4,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "TestLogger_Inform() -> Hello World!\n",
		},
	}

	for testNumber, test := range tests {

		var output strings.Builder
		var logger Logger
		logger.Writer = &output
		logger = logger.Level(test.level)
		logger.Inform(test.Parameters...)

		if expected, actual := test.Expected, output.String(); expected != actual {
			t.Errorf("For test #%d, Inform did not log what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

	}
}

func TestLogger_Log(t *testing.T) {

	tests := []struct {
		level      uint8
		Parameters []interface{}
		Expected   string
	}{
		{
			level:      0,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      1,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      2,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      3,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      4,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      5,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "TestLogger_Log() -> Hello World!\n",
		},
	}

	for testNumber, test := range tests {

		var output strings.Builder
		var logger Logger
		logger.Writer = &output
		logger = logger.Level(test.level)
		logger.Log(test.Parameters...)

		if expected, actual := test.Expected, output.String(); expected != actual {
			t.Errorf("For test #%d, Log did not log what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

	}
}

func TestLogger_Trace(t *testing.T) {

	tests := []struct {
		level      uint8
		Parameters []interface{}
		Expected   string
	}{
		{
			level:      0,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      1,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      2,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      3,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      4,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      5,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "",
		},
		{
			level:      6,
			Parameters: []interface{}{"Hello World!"},
			Expected:   "TestLogger_Trace() -> Hello World!\n",
		},
	}

	for testNumber, test := range tests {

		var output strings.Builder
		var logger Logger
		logger.Writer = &output
		logger = logger.Level(test.level)
		logger.Trace(test.Parameters...)

		if expected, actual := test.Expected, output.String(); expected != actual {
			t.Errorf("For test #%d, Trace did not log what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

	}
}