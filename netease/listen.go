package netease

import (
	"encoding/json"
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"log"
	"net/http"
	"nowPlaying/config"
	"nowPlaying/httpserver"
)

//var song *Netease
var song json.RawMessage

func Listen() {
	cfg := config.Netease
	if cfg == nil || cfg.Enable == false || cfg.ListenDir == "" {
		return
	}
	http.HandleFunc("/netease/nowPlaying", nowPlaying)
	http.HandleFunc("/netease/playQueue", playQueue)
	http.HandleFunc("/netease/song", songUrl)

	updateNeteaseMusic(cfg.ListenDir)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	err = watcher.Add(cfg.ListenDir)
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
				updateNeteaseMusic(cfg.ListenDir)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}

func updateNeteaseMusic(dir string) {
	content, err := ioutil.ReadFile(dir + "history")
	if err != nil {
		return
	}
	var neteases []json.RawMessage

	err = json.Unmarshal(content, &neteases)

	if err != nil {
		log.Fatalln(err)
	}

	if len(neteases) > 0 {
		song = neteases[0]
		var netease Netease
		err := json.Unmarshal(song, &netease)

		if err == nil {
			updateGithubStatus(netease)
		}
	}
}

func nowPlaying(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Content-Type", "application/json")

	_, _ = writer.Write(song)
}

func playQueue(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Content-Type", "application/json")

	content, err := ioutil.ReadFile(config.Netease.ListenDir + "queue")
	if err != nil {
		httpserver.FailedResponse(err, writer)
		return
	}
	_, _ = writer.Write(content)
}

func songUrl (writer http.ResponseWriter, _ *http.Request) {
//
}
