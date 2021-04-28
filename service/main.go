package service

import (
	"errors"
	"golang.org/x/sys/windows/svc/mgr"
	"nowPlaying/service/windows"
	"runtime"
)
var err = errors.New("unsupported operating system" + runtime.GOOS)

func Install(path, name, desc string, param ...string) error {
	if runtime.GOOS == "windows" {
		return windows.Install(path, name, desc, param...)
	}

	return err
}

func Uninstall(name string) error {
	if runtime.GOOS == "windows" {
		return windows.Uninstall(name)
	}

	return err
}

func Start(name string) error {
	if runtime.GOOS == "windows" {
		return windows.Start(name)
	}

	return err
}

func Stop(name string) error {
	if runtime.GOOS == "windows" {
		return windows.Stop(name)
	}

	return err
}

func InServiceModel() bool {
	if runtime.GOOS == "windows" {
		return windows.InServiceModel()
	}

	return false
}

func Installed(name string) (bool, error) {
	mg, err := mgr.Connect()

	if err != nil {
		return false, err
	}
	defer mg.Disconnect()

	service, err := mg.OpenService(name)

	if err != nil {
		return false, nil
	}
	_ = service.Close()

	return true, nil
}

func RunAsService(name string, start, stop func()) error {
	if runtime.GOOS == "windows" {
		return windows.RunAsService(name, start, stop)
	}

	return err
}