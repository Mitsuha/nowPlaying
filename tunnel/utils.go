package tunnel

import (
	"nowPlaying/config"
)

func privateKetPath() string {
	return config.App.HomeDir + "\\.ssh\\id_rsa"
}
