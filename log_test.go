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
