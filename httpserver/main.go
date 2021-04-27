package httpserver

import (
	"encoding/json"
	"log"
	"net/http"
)

func Start(addr string) {
	http.HandleFunc("/", func(writer http.ResponseWriter, _ *http.Request) {
		_, _ = JsonResponse(map[string]string{
			"version": "dev version",
		}, writer)
	})

	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatalln(err)
	}
}

func FailedResponse(err error) []byte {
	response, _ := json.Marshal(map[string]string{
		"status":  "error",
		"message": err.Error(),
	})
	return response
}

func JsonResponse(data interface{}, writer http.ResponseWriter) (int,error) {
	response, err := json.Marshal(map[string]string{
		"version": "dev version",
	})
	if err != nil {
		return 0, err
	}

	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Content-Type", "application/json")

	return writer.Write(response)
}