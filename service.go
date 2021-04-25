package main

import (
	"log"
	"nowPlaying/config"
	nos "nowPlaying/os"
	"nowPlaying/service"
	"os"
)

func registerService() {
	if *flagServiceInstall {
		ap, err := nos.AppPath()
		if err != nil {
			log.Fatalln(err)
		}

		err = service.Install(ap, config.ServiceName, config.ServiceDescription)

		if err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}

	if *flagServiceUninstall {
		err := service.Uninstall(config.ServiceName)

		if err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}

	if *flagServiceStop {
		err := service.Stop(config.ServiceName)

		if err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}

	if *flagServiceStart {
		err := service.Start(config.ServiceName)

		if err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}
}