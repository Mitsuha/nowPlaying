package tunnel

import (
	"nowPlaying/os"
)

func privateKetPath() string {
	return os.HomeDirectory() + "\\.ssh\\id_rsa"
}
