package main

import (
	"log"
	"nowPlaying/config"
	"nowPlaying/httpserver"
	"nowPlaying/netease"
	"nowPlaying/service"
	"nowPlaying/tunnel"
	"os"
)

func init() {
	//fmt.Println(time.Now().Format("2006-01-02T15:04:05Z07:00"))
	//os.Exit(1)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config.Initialization()

	logf, err := os.OpenFile(config.Shell.LogPath, os.O_WRONLY, 0)
	if err != nil {
		log.Fatalln(err)
	}
	if service.InServiceModel() {
		log.SetOutput(logf)
	}

}

func main() {
	registerService()

	if service.InServiceModel() {
		err := service.RunAsService(config.ServiceName, start, stop)

		if err != nil {
			log.Fatalln(err)
		}
	}else{
		start()
	}
}

func start() {
	go netease.Listen()
	go httpserver.Start()

	if config.Tunnel.Enable {
		if config.Tunnel.Type == "ssh" {
			sshTunnel, err := tunnel.Ssh(config.Tunnel)
			if err != nil {
				log.Fatalln(err)
			}

			err = sshTunnel.MappingRemote()
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	select {}
}

func stop() {

}