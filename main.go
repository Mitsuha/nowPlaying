package main

import (
	"flag"
	"log"
	"nowPlaying/config"
	"nowPlaying/httpserver"
	"nowPlaying/models"
	"nowPlaying/netease"
	"nowPlaying/service"
	"nowPlaying/tunnel"
	"os"
)

var (
    flagServiceInstall   = flag.Bool("install", false, "Install service")
    flagServiceUninstall = flag.Bool("uninstall", false, "Remove service")
    flagServiceStart     = flag.Bool("start", false, "Start service")
    flagServiceStop      = flag.Bool("stop", false, "Stop service")
)

func init() {
	flag.Parse()

	logf, err := os.OpenFile("C:\\Users\\hhx\\Desktop\\nowPlaying.log", os.O_WRONLY, 0)
	if err != nil {
		log.Fatalln(err)
	}
	if service.InServiceModel() {
		log.SetOutput(logf)
	}

    log.SetFlags(log.LstdFlags | log.Lshortfile)

    config.InitConfig()
}

func main() {
	registerService()

	if service.InServiceModel() {
		err := service.RunAsService("nowPlaying", application, func() {})

		if err != nil {
			log.Fatalln(err)
		}
	}else{
		application()
	}
}

func application() {
	var nowPlaying models.Netease
	go netease.Listen(&nowPlaying)
	go httpserver.Start(&nowPlaying, config.LocalAddr)

	sshTunnel, err := tunnel.Ssh(config.RemoteServer, "正确的密码", config.LocalAddr, config.RemoteListen)
	if err != nil {
		log.Fatalln(err)
	}

	err = sshTunnel.MappingRemote()
	if err != nil {
		log.Fatalln(err)
	}

	select {}
}