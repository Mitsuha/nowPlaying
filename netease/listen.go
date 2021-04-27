package netease

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"log"
	"net/http"
	"nowPlaying/config"
	"nowPlaying/httpserver"
)

//var song *Netease
var song json.RawMessage
var cfg *config.NeteaseCfg

func Listen(neteaseCfg *config.NeteaseCfg) {
	cfg = neteaseCfg
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
		fmt.Println(string(neteases[0]))

		song = neteases[0]
	}
}

func nowPlaying(writer http.ResponseWriter, _ *http.Request) {
	//_, err := httpserver.JsonResponse(song, writer)
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Content-Type", "application/json")

	_, _ = writer.Write(song)
	//if err != nil {
	//	log.Println(err)
	//}
}

func playQueue(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Content-Type", "application/json")

	content, err := ioutil.ReadFile(cfg.ListenDir + "queue")
	if err != nil {
		_, _ = writer.Write(httpserver.FailedResponse(err))
		return
	}
	_, _ = writer.Write(content)
}

func songUrl (writer http.ResponseWriter, _ *http.Request) {
//
}
