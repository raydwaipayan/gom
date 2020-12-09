package process

import (
	"os"
	"path/filepath"
	"testing"
)

var proc Proc

func init() {
	wd, _ := os.Getwd()
	dir := filepath.Join(wd, "../test")
	proc = Proc{
		Name: "test_program.go",
		Cmd:  "/usr/local/go/bin/go",
		Argv: []string{"go", "run", filepath.Join(dir, "test_program.go")},
		Path: dir,
	}
}

func TestStart(t *testing.T) {
	err := proc.Start()
	if err != nil {
		t.Errorf("Could not start process: %s", err)
	}
}

func TestRestart(t *testing.T) {
	err := proc.Restart()
	if err != nil {
		t.Errorf("Could not restart process")
	}
}

func TestStop(t *testing.T) {
	err := proc.Stop()
	if err != nil {
		t.Errorf("Could not stop process")
	}
}
