package log

import (
	"strings"
	"testing"
)

func TestLogger_Begin(t *testing.T) {
	var output strings.Builder
	var logger Logger
	logger.Writer = &output
	logger.Begin()
	if output.String() != "BEGIN\n" {
		t.Errorf("Expected BEGIN received '%s'", output.String())
	}
}

func TestLogger_End(t *testing.T) {
	var output strings.Builder
	var logger Logger
	logger.Writer = &output
	logger.End()
	if output.String() != "END\n" {
		t.Errorf("Expected END received '%s'", output.String())
	}
}

func TestLogger_Logf(t *testing.T) {
	var output strings.Builder
	var logger Logger
	logger.Writer = &output
	logger.Logf("Hello World!")
	if output.String() != "Hello World!\n" {
		t.Errorf("Expected Hello World! received '%s'", output.String())
	}
}