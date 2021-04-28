package netease

//type Array struct {
//	json.RawMessage
//}

type Netease struct {
	Track struct {
		Album struct {
			Name       string        `json:"name"`
			Picurl     string        `json:"picUrl"`
			//Alias      []interface{} `json:"alias"`
			//Transnames []interface{} `json:"transNames"`
		} `json:"album"`
		//Alias   []string `json:"alias"`
		Artists []struct {
			ID    int           `json:"id"`
			Name  string        `json:"name"`
			//Tns   []interface{} `json:"tns"`
			//Alias []interface{} `json:"alias"`
		} `json:"artists"`
		Duration    int    `json:"duration"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
	} `json:"track"`
	Text         string `json:"text"`
	Nickname     string `json:"nickName"`
	Startlogtime int64  `json:"startlogtime"`
	Time         int64  `json:"time"`
	LastPlayInfo struct{
		RetJson struct{
			Url string `json:"url"`
		} `json:"retJson"`
	} `json:"lastPlayInfo"`
}
