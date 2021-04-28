package netease

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"nowPlaying/config"
	"strings"
	"time"
)

func updateGithubStatus(netease Netease) {
	client := http.Client{
		Timeout:       30*time.Second,
	}

	request, _ := http.NewRequest("GET", "https://github.com/" + config.Github.Username, nil)
	request.Header.Set("cookie", config.Github.Cookies)
	response, err := client.Do(request)
	if err != nil {
		return
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return
	}
	inputs := document.Find(".js-user-status-form").Find("input[name='authenticity_token']")

	token, exists := inputs.Attr("value")
	if ! exists {
		return
	}

	var songName = netease.Track.Name
	if len(songName) < 10 {
		songName += "-" + netease.Track.Artists[0].Name
	}
	request, _ = http.NewRequest("POST", "https://github.com/users/status", strings.NewReader(url.Values{
		"_method": []string{"PUT"},
		"authenticity_token": []string{token},
		"emoji": []string{":musical_note:"},
		"expires_at": []string{time.Now().Add(4 * time.Minute).Format(time.RFC3339)},
		"message": []string{"Listening: " + songName + " (Real-time Data)"},
	}.Encode()))

	request.Header.Set("cookie", config.Github.Cookies)
	response, _ = client.Do(request)
	//body, _ := ioutil.ReadAll(response.Body)

	//fmt.Println(string(body))
	_ = response.Body.Close()
}
