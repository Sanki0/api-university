package app

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
    var application App
    application.Initialize()
	code := m.Run()
	os.Exit(code)
}