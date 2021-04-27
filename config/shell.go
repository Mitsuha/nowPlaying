package config

import "flag"

type ShellCfg struct {
	Install, Uninstall, Start, Stop bool
	IniPath string
}

func newShell() *ShellCfg {
	var (
		install   = flag.Bool("install", false, "Install service")
		uninstall = flag.Bool("uninstall", false, "Remove service")
		start     = flag.Bool("start", false, "Start service")
		stop      = flag.Bool("stop", false, "Stop service")
		iniPath   = flag.String("f", "", "Stop service")
	)

	flag.Parse()

	return &ShellCfg{
		Install:   *install,
		Uninstall: *uninstall,
		Start:     *start,
		Stop:      *stop,
		IniPath:   *iniPath,
	}
}