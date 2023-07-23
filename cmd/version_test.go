package cmd

import (
	"runtime/debug"
	"testing"
)

func TestGetVersion(t *testing.T) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		t.Fatal("Error: no debug info")
	}
	t.Log(info.Main.Version)
}
