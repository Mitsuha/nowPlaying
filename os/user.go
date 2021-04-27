package os

import (
	"fmt"
	"golang.org/x/sys/windows/svc/mgr"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
)

func HomeDirectory() string {
	home, err := os.UserHomeDir()

	if err != nil {
		log.Fatalln(err)
	}
	return home
}

func Username() string {
	u, err := user.Current()

	if err != nil {
		log.Fatalln(err)
	}

	return u.Name
}

func AppPath() (string, error) {
	prog := os.Args[0]
	p, err := filepath.Abs(prog)
	if err != nil {
		return "", err
	}
	fi, err := os.Stat(p)
	if err == nil {
		if !fi.Mode().IsDir() {
			return p, nil
		}
		err = fmt.Errorf("winsvc.GetAppPath: %s is directory", p)
	}
	if filepath.Ext(p) == "" {
		p += ".exe"
		fi, err := os.Stat(p)
		if err == nil {
			if !fi.Mode().IsDir() {
				return p, nil
			}
			err = fmt.Errorf("winsvc.GetAppPath: %s is directory", p)
		}
	}
	return "", err

}
func HasHeightPermission() bool {
	if runtime.GOOS == "windows" {
		m, err := mgr.Connect()

		if err != nil {
			return false
		}
		_ = m.Disconnect()
		
		return true
	}
	if runtime.GOOS == "linux" {
		return Username() == "root"
	}
	return false
}

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