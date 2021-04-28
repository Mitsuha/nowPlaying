package httpserver

import (
	"encoding/json"
	"log"
	"net/http"
	"nowPlaying/config"
)

func Start() {
	http.HandleFunc("/", func(writer http.ResponseWriter, _ *http.Request) {
		_, _ = JsonResponse(map[string]string{
			"version": "dev version",
		}, writer)
	})

	err := http.ListenAndServe(config.App.Listen, nil)

	if err != nil {
		log.Fatalln(err)
	}
}

func FailedResponse(err error, writer http.ResponseWriter) {
	_, _ = JsonResponse(map[string]string{
		"status":  "error",
		"message": err.Error(),
	}, writer)
}

func JsonResponse(data interface{}, writer http.ResponseWriter) (int,error) {
	response, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Content-Type", "application/json")

	return writer.Write(response)
}