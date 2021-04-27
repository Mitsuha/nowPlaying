package main

import (
	"fmt"
	"log"
	"nowPlaying/config"
	"nowPlaying/httpserver"
	"nowPlaying/netease"
	"nowPlaying/service"
	"nowPlaying/tunnel"
	"os"
)

func init() {
	config.Initialization()

	logf, err := os.OpenFile("C:\\Users\\hhx\\Desktop\\nowPlaying.log", os.O_WRONLY, 0)
	if err != nil {
		log.Fatalln(err)
	}
	if service.InServiceModel() {
		log.SetOutput(logf)
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	registerService()

	if service.InServiceModel() {
		err := service.RunAsService("nowPlaying", app, func() {})

		if err != nil {
			log.Fatalln(err)
		}
	}else{
		app()
	}
}

func app() {
	go netease.Listen(config.Netease)
	//time.Sleep(1*time.Second)
	go httpserver.Start(config.App.Listen)

	if config.Tunnel.Enable {
		if config.Tunnel.Type == "ssh" {
			fmt.Println(config.Tunnel)
			//cfg := config.Tunnel
			//sshTunnel, err := tunnel.Ssh(cfg.User, cfg.Host, cfg.Password, config.App.Listen, "0.0.0.0:" + cfg.Port)
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