package config

import (
	"gopkg.in/ini.v1"
	"log"
	nos "nowPlaying/os"
	"os"
)

//var Config *ini.File
//
//var LocalAddr string
//var RemoteServer string
//var RemoteListen string
//var flagConfigPath = flag.String("f", "", "ini file")
//var UserNetEaseMusicFilePath = ""

var App *AppCfg
var Shell *ShellCfg
var Netease *NeteaseCfg
var Tunnel *TunnelCfg
var Config *ini.File

func Initialization() {
	Shell = newShell()
	var iniPath string

	if Shell.IniPath == "" {
		iniPath = nos.HomeDirectory() + "\\.nowPlaying\\config.ini"
	}

	Config, err := ini.Load(iniPath)
	if err != nil {
		log.Fatalln(err)
	}

	App = &AppCfg{
		Listen:  Config.Section("app").Key("listen").String(),
		IniPath: Shell.IniPath,
		HomeDir: nos.HomeDirectory(),
	}

	enableTunnel, _ := Config.Section("netease").Key("enable").Bool()
	Tunnel = &TunnelCfg{
		Enable:  enableTunnel,
		Type:    Config.Section("tunnel").Key("type").String(),
		User:    Config.Section("tunnel").Key("user").String(),
		Host:    Config.Section("tunnel").Key("host").String(),
		Port:    Config.Section("tunnel").Key("port").String(),
		Forward: App.Listen,
	}

	enableNetease, _ := Config.Section("netease").Key("enable").Bool()
	neteaseListenDir := Config.Section("netease").Key("ListenDir").String()

	if neteaseListenDir == "" {neteaseListenDir = App.HomeDir + WindowsNetEaseMusicFilePath}
	Netease = &NeteaseCfg{
		Enable: enableNetease ,
		ListenDir: neteaseListenDir,
	}

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
