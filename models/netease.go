package models

type Netease struct {
	Track struct {
		Album struct {
			Name       string        `json:"name"`
			Picurl     string        `json:"picUrl"`
			Alias      []interface{} `json:"alias"`
			Transnames []interface{} `json:"transNames"`
		} `json:"album"`
		Alias   []string `json:"alias"`
		Artists []struct {
			ID    int           `json:"id"`
			Name  string        `json:"name"`
			Tns   []interface{} `json:"tns"`
			Alias []interface{} `json:"alias"`
		} `json:"artists"`
		Copyrightid int    `json:"copyrightId"`
		Duration    int    `json:"duration"`
		ID          int    `json:"id"`
		Mvid        int    `json:"mvid"`
		Name        string `json:"name"`
		Hmusic      struct {
			Bitrate     int `json:"bitrate"`
			Dfsid       int `json:"dfsId"`
			Size        int `json:"size"`
			Volumedelta int `json:"volumeDelta"`
		} `json:"hMusic"`
		Mmusic struct {
			Bitrate     int `json:"bitrate"`
			Dfsid       int `json:"dfsId"`
			Size        int `json:"size"`
			Volumedelta int `json:"volumeDelta"`
		} `json:"mMusic"`
		Lmusic struct {
			Bitrate     int `json:"bitrate"`
			Dfsid       int `json:"dfsId"`
			Size        int `json:"size"`
			Volumedelta int `json:"volumeDelta"`
		} `json:"lMusic"`
	} `json:"track"`
	ID           string `json:"id"`
	Tid          int    `json:"tid"`
	Href         string `json:"href"`
	Text         string `json:"text"`
	Nickname     string `json:"nickName"`
	Startlogtime int64  `json:"startlogtime"`
	Time         int64  `json:"time"`
}
