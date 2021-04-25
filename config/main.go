package config

import (
	"flag"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"nowPlaying/common"
	nos "nowPlaying/os"
	"os"
)

var Config *ini.File

var LocalAddr string
var RemoteServer string
var RemoteListen string
var flagConfigPath = flag.String("f", "", "ini file")

func InitConfig() {
	//fmt.Println("ffff",*flagConfigPath)
	log.Println(nos.HomeDirectory())
	var iniPath = nos.HomeDirectory() + "\\.nowPlaying"

	var err error
	if ! common.FileExists(iniPath) {
		err = createConfigFile(iniPath)
		if err != nil {
			log.Fatalln(err)
		}

		err = common.OpenFilExplorer(iniPath)
		if err != nil {
			fmt.Println(err)
			fmt.Println("First running, please edit your configuration file on this path " + iniPath)
		}
		fmt.Println("First running, please edit your configuration file")
		os.Exit(1)
	}

	Config, err = ini.Load(iniPath + "\\config.ini")

	if err != nil {
		log.Fatalln(err)
	}

	LocalAddr = Config.Section("web").Key("addr").String()
	RemoteServer = Config.Section("tunnel").Key("server").String()
	RemoteListen = Config.Section("tunnel").Key("listen").String()
}

func createConfigFile(iniPath string) error {
	err := os.Mkdir(iniPath, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create(iniPath + "\\config.ini")

	if err != nil {
		return err
	}
	return f.Close()
}
