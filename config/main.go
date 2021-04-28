package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	nos "nowPlaying/os"
	"os"
	"path/filepath"
)

var App *AppCfg
var Shell *ShellCfg
var Netease *NeteaseCfg
var Tunnel *TunnelCfg
var Github *GithubCfg
var Config *ini.File

func Initialization() {
	Shell = newShell()
	if ! nos.FileExists(Shell.IniPath) {
		file, err := createFile(Shell.IniPath)
		if err != nil {
			log.Fatalln(err)
		}
		_ = writeConfigExample(file)
		_ = nos.OpenFilExplorer(filepath.Dir(Shell.IniPath))
		fmt.Println("Please edit your configuration file before starting: " + Shell.IniPath)
	}

	if ! nos.FileExists(Shell.LogPath) {
		file, err := createFile(Shell.LogPath)
		if err != nil {
			log.Fatalln(err)
		}
		file.Close()
	}

	Config, err := ini.LoadSources(ini.LoadOptions{SpaceBeforeInlineComment: true}, Shell.IniPath)
	if err != nil {
		log.Fatalln(err)
	}

	App = &AppCfg{
		Listen:  Config.Section("app").Key("listen").String(),
		IniPath: Shell.IniPath,
		HomeDir: Shell.HomeDir,
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

	enableGithub, _ := Config.Section("netease").Key("enable").Bool()
	Github = &GithubCfg{
		Enable: enableGithub ,
		Cookies: Config.Section("github").Key("cookies").String(),
		Username: Config.Section("github").Key("username").String(),
	}
	if Github.Cookies == "" {Github.Enable = false}
}

func createFile(file string) (*os.File, error) {
	path := filepath.Dir(file)
	if ! nos.FileExists(path) {
		err := os.MkdirAll(filepath.Dir(file), 0666)

		if err != nil {
			log.Fatalln(err)
		}
	}

	return os.Create(file)
}

func writeConfigExample(file *os.File) error {
	_, err := file.Write([]byte(`[app]
listen = 0.0.0.0:5000

[tunnel]
enable = true
type = ssh
user = 
password = 
host = 
port = 

[netease]
enable = true

[github]
enable = true
username = 
cookies = ""
`))
	return err
}