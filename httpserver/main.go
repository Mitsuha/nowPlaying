package httpserver

import (
	"encoding/json"
	"log"
	"net/http"
	"nowPlaying/models"
)

func Start(nowPlaying *models.Netease, addr string) {
	http.HandleFunc("/", func(writer http.ResponseWriter, _ *http.Request) {
		response, err := json.Marshal(nowPlaying)

		if err != nil {
			response, _ = json.Marshal(map[string]string{
				"status":  "error",
				"message": err.Error(),
			})
		}
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Content-Type", "application/json")
		_, _ = writer.Write(response)
	})

	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatalln(err)
	}
}
