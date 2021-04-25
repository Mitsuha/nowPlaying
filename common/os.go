package common

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func OpenFilExplorer(path string) error {
	var command = map[string]string{
		"windows": "explorer",
	}

	if cmd, ok := command[runtime.GOOS]; ok {
		return exec.Command(cmd, path).Start()
	}
	return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
}