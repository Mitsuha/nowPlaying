package config

import (
	"flag"
	nos "nowPlaying/os"
)

type ShellCfg struct {
	Install, Uninstall, Start, Stop bool
	HomeDir string
	IniPath string
	LogPath string
}

func newShell() *ShellCfg {
	var (
		install   = flag.Bool("install", false, "Install service")
		uninstall = flag.Bool("uninstall", false, "Remove service")
		start     = flag.Bool("start", false, "Start service")
		stop      = flag.Bool("stop", false, "Stop service")
		homeDir   = flag.String("h", "", "User's home directory")
		iniPath   = flag.String("f", "", "Configuration file path")
		logPath   = flag.String("l", "", "Log file path")
	)

	flag.Parse()

	if *homeDir == "" {
		*homeDir = nos.HomeDirectory()
	}

	if *iniPath == "" {
		*iniPath = *homeDir + "\\.nowPlaying\\config.ini"
	}

	if *logPath == "" {
		*logPath = *homeDir + "\\.nowPlaying\\app.log"
	}

	return &ShellCfg{
		Install:   *install,
		Uninstall: *uninstall,
		Start:     *start,
		Stop:      *stop,
		HomeDir:   *homeDir,
		IniPath:   *iniPath,
		LogPath:   *logPath,
	}
}