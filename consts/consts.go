package consts

import (
	"nowPlaying/common"
	"nowPlaying/os"
)

const (
	WindowsNetEaseMusicFilePath = "\\AppData\\Local\\Netease\\CloudMusic\\webdata\\file\\history"
)

var UserNetEaseMusicFilePath = ""

func init() {
	netEase := os.HomeDirectory() + WindowsNetEaseMusicFilePath

	if common.FileExists(netEase) {
		UserNetEaseMusicFilePath = netEase
	}
}
