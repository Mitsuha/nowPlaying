package netease

import (
	"encoding/json"
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"log"
	"nowPlaying/consts"
	"nowPlaying/models"
)

func Listen(nowPlaying *models.Netease) {
	if consts.UserNetEaseMusicFilePath == "" {
		return
	}
	updateNeteaseMusic(nowPlaying)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	err = watcher.Add(consts.UserNetEaseMusicFilePath)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if event.Op&fsnotify.Remove == fsnotify.Remove {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				updateNeteaseMusic(nowPlaying)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}

func updateNeteaseMusic(nowPlaying *models.Netease) {
	content, err := ioutil.ReadFile(consts.UserNetEaseMusicFilePath)

	if err != nil {
		return
	}
	var neteases []models.Netease

	err = json.Unmarshal(content, &neteases)

	if err != nil {
		log.Fatalln(err)
	}

	if len(neteases) > 0 {
		*nowPlaying = neteases[0]
	}
}
