package common

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Comment / un comment below to see log output while testing
	// log.ConfigureLogger(log.WithNullLogger())
	// log.ConfigureLogger(log.WithLevel("debug"))
	os.Exit(m.Run())
}
