package main

import (
	"log"
	"nowPlaying/config"
	nos "nowPlaying/os"
	"nowPlaying/service"
	"os"
)

func registerService() {
	if config.Shell.Install {
		ap, err := nos.AppPath()
		if err != nil {
			log.Fatalln(err)
		}

		//err = service.Install(ap, config.ServiceName, config.ServiceDescription, fmt.Sprintf(
		//	"-f %s -h %s",
		//	config.App.IniPath, config.App.HomeDir,
		//	))
		err = service.Install(ap, config.ServiceName, config.ServiceDescription, "-f", config.App.IniPath, "-h", config.App.HomeDir)

		if err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}

	if config.Shell.Uninstall {
		err := service.Uninstall(config.ServiceName)

		if err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}

	if config.Shell.Start {
		err := service.Start(config.ServiceName)

		if err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}

	if config.Shell.Stop {
		err := service.Stop(config.ServiceName)

		if err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}
}